package service

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go-wind-uba/app/core/service/internal/data"
	"go-wind-uba/app/core/service/internal/data/clickhouse"
	dorisData "go-wind-uba/app/core/service/internal/data/doris"
	"go-wind-uba/app/core/service/internal/data/ent"

	internalMessageV1 "go-wind-uba/api/gen/go/internal_message/service/v1"
	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

const (
	reportScanInterval = time.Minute
	reportLockKey      = "uba:lock:report-engine"
	reportLockTTL      = 55 * time.Second

	// reportEventType 报告生成的 webhook 事件类型，与 Webhook.event_types 配对。
	reportEventType = "report.generated"

	// reportMinInterval 最小执行间隔（评估粒度按分钟，但报告别配得太密）。
	reportMinInterval = time.Hour

	// reportCSVMaxRows CSV 附件的最大行数（防止超长消息）。
	reportCSVMaxRows = 2000
)

// reportAnalytics 报告执行所需的分析查询能力（Doris/ClickHouse 双实现共用签名）。
type reportAnalytics interface {
	EventTrend(ctx context.Context, req *ubaV1.EventTrendRequest) (*ubaV1.EventTrendResponse, error)
	ActiveUsers(ctx context.Context, req *ubaV1.ActiveUsersRequest) (*ubaV1.ActiveUsersResponse, error)
	GroupBy(ctx context.Context, req *ubaV1.GroupByRequest) (*ubaV1.GroupByResponse, error)
	Funnel(ctx context.Context, req *ubaV1.FunnelRequest) (*ubaV1.FunnelResponse, error)
	Retention(ctx context.Context, req *ubaV1.RetentionRequest) (*ubaV1.RetentionResponse, error)
}

// ReportEngine 定时报告执行器：按订阅的间隔重跑分析查询，把结果摘要
// （文本 + CSV）推送到站内信/网络钩子。到期判定用 last_run_at + interval_minutes。
type ReportEngine struct {
	log *log.Helper

	reportRepo    *data.ReportSubscriptionRepo
	msgRepo       *data.InternalMessageRepo
	recipientRepo *data.InternalMessageRecipientRepo
	webhookSender *WebhookSender
	redis         *redis.Client

	analytics   reportAnalytics
	chAnalytics reportAnalytics

	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewReportEngine(
	ctx *bootstrap.Context,
	reportRepo *data.ReportSubscriptionRepo,
	msgRepo *data.InternalMessageRepo,
	recipientRepo *data.InternalMessageRecipientRepo,
	webhookSender *WebhookSender,
	redisClient *redis.Client,
	analytics *dorisData.AnalyticsRepo,
	chAnalytics *clickhouse.AnalyticsRepo,
) *ReportEngine {
	return &ReportEngine{
		log:           ctx.NewLoggerHelper("report-engine/service/core-service"),
		reportRepo:    reportRepo,
		msgRepo:       msgRepo,
		recipientRepo: recipientRepo,
		webhookSender: webhookSender,
		redis:         redisClient,
		analytics:     analytics,
		chAnalytics:   chAnalytics,
	}
}

// Start 实现 transport.Server：启动执行循环。
func (s *ReportEngine) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	s.wg.Add(1)
	go s.loop(runCtx)

	return nil
}

// Stop 实现 transport.Server：停止执行循环。
func (s *ReportEngine) Stop(_ context.Context) error {
	if s.cancel != nil {
		s.cancel()
	}
	s.wg.Wait()
	return nil
}

func (s *ReportEngine) loop(ctx context.Context) {
	defer s.wg.Done()

	ticker := time.NewTicker(reportScanInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.tick(ctx)
		}
	}
}

func (s *ReportEngine) tick(ctx context.Context) {
	if s.redis != nil {
		acquired, err := s.redis.SetNX(ctx, reportLockKey, time.Now().UnixMilli(), reportLockTTL).Result()
		if err != nil {
			s.log.Errorf("acquire report-engine lock failed: %s", err.Error())
			return
		}
		if !acquired {
			return
		}
	}

	due, err := s.reportRepo.ListDueSubscriptions(ctx, time.Now())
	if err != nil {
		s.log.Errorf("list due report subscriptions failed: %s", err.Error())
		return
	}

	now := time.Now()
	for _, sub := range due {
		summary, csv, runErr := s.runSubscription(ctx, sub, now)
		if runErr != nil {
			s.log.Errorf("run report subscription [%d] failed: %s", sub.ID, runErr.Error())
			continue
		}
		s.deliver(ctx, sub, summary, csv, now)
		if err := s.reportRepo.UpdateRunState(ctx, sub.ID, now); err != nil {
			s.log.Errorf("record report run state failed: %s", err.Error())
		}
	}
}

// runSubscription 执行一个订阅对应的分析查询，返回 (摘要文本, CSV 内容)。
func (s *ReportEngine) runSubscription(ctx context.Context, sub *ent.ReportSubscription, now time.Time) (string, string, error) {
	impl := s.analyticsImpl()

	interval := time.Duration(sub.IntervalMinutes) * time.Minute
	if interval < reportMinInterval {
		interval = reportMinInterval
	}

	switch derefStr(sub.Model) {
	case "event-trend":
		req := &ubaV1.EventTrendRequest{}
		if err := parseReportParams(derefStr(sub.ParamsJSON), req); err != nil {
			return "", "", err
		}
		req.TimeRange = ensureReportRange(req.TimeRange, interval, now)

		resp, err := impl.EventTrend(ctx, req)
		if err != nil {
			return "", "", err
		}

		points := resp.GetPoints()
		latest := 0.0
		if len(points) > 0 {
			latest = points[len(points)-1].GetValue()
		}
		summary := fmt.Sprintf("区间事件总量 %d，最新桶 %.0f（%d 个数据点）。", resp.GetTotal(), latest, len(points))

		headers := []string{"timestamp", "value"}
		rows := make([][]string, 0, len(points))
		for _, p := range points {
			rows = append(rows, []string{
				fmt.Sprintf("%d", p.GetTimestamp()),
				fmt.Sprintf("%.0f", p.GetValue()),
			})
		}
		return summary, csvFromRows(headers, rows), nil

	case "active-users":
		req := &ubaV1.ActiveUsersRequest{}
		if err := parseReportParams(derefStr(sub.ParamsJSON), req); err != nil {
			return "", "", err
		}
		req.TimeRange = ensureReportRange(req.TimeRange, interval, now)

		resp, err := impl.ActiveUsers(ctx, req)
		if err != nil {
			return "", "", err
		}

		points := resp.GetPoints()
		wau, mau := int64(0), int64(0)
		if len(points) > 0 {
			wau = points[len(points)-1].GetWau()
			mau = points[len(points)-1].GetMau()
		}
		summary := fmt.Sprintf("最新 DAU %d，WAU %d，MAU %d。", resp.GetLatestDau(), wau, mau)

		headers := []string{"timestamp", "dau", "wau", "mau"}
		rows := make([][]string, 0, len(points))
		for _, p := range points {
			rows = append(rows, []string{
				fmt.Sprintf("%d", p.GetTimestamp()),
				fmt.Sprintf("%d", p.GetDau()),
				fmt.Sprintf("%d", p.GetWau()),
				fmt.Sprintf("%d", p.GetMau()),
			})
		}
		return summary, csvFromRows(headers, rows), nil

	case "group-by":
		req := &ubaV1.GroupByRequest{}
		if err := parseReportParams(derefStr(sub.ParamsJSON), req); err != nil {
			return "", "", err
		}
		req.TimeRange = ensureReportRange(req.TimeRange, interval, now)

		resp, err := impl.GroupBy(ctx, req)
		if err != nil {
			return "", "", err
		}

		buckets := resp.GetBuckets()
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("维度 %s 共 %d 个取值，Top：", resp.GetDimension(), len(buckets)))
		for i, b := range buckets {
			if i >= 5 {
				break
			}
			if i > 0 {
				sb.WriteString("、")
			}
			sb.WriteString(fmt.Sprintf("%s(%v)", b.GetLabel(), b.GetValue()))
		}
		summary := sb.String()

		headers := []string{"label", "value"}
		rows := make([][]string, 0, len(buckets))
		for _, b := range buckets {
			rows = append(rows, []string{b.GetLabel(), fmt.Sprintf("%v", b.GetValue())})
		}
		return summary, csvFromRows(headers, rows), nil

	case "funnel":
		req := &ubaV1.FunnelRequest{}
		if err := parseReportParams(derefStr(sub.ParamsJSON), req); err != nil {
			return "", "", err
		}
		req.TimeRange = ensureReportRange(req.TimeRange, interval, now)

		resp, err := impl.Funnel(ctx, req)
		if err != nil {
			return "", "", err
		}

		summary := fmt.Sprintf("进入 %d 人，完成全部步骤 %d 人，整体转化率 %.2f%%。",
			resp.GetEnteredUsers(), resp.GetCompletedUsers(), resp.GetOverallConversion()*100)

		steps := resp.GetSteps()
		headers := []string{"step", "event", "count", "overall_rate"}
		rows := make([][]string, 0, len(steps))
		for _, st := range steps {
			rows = append(rows, []string{
				fmt.Sprintf("%d", st.GetStepIndex()),
				st.GetEventName(),
				fmt.Sprintf("%d", st.GetCount()),
				fmt.Sprintf("%.4f", st.GetOverallRate()),
			})
		}
		return summary, csvFromRows(headers, rows), nil

	case "retention":
		req := &ubaV1.RetentionRequest{}
		if err := parseReportParams(derefStr(sub.ParamsJSON), req); err != nil {
			return "", "", err
		}
		req.TimeRange = ensureReportRange(req.TimeRange, interval, now)

		resp, err := impl.Retention(ctx, req)
		if err != nil {
			return "", "", err
		}

		cohorts := resp.GetCohorts()
		day1Sum, day1Count := 0.0, 0
		for _, c := range cohorts {
			for _, cell := range c.GetCells() {
				if cell.GetOffsetDays() == 1 {
					day1Sum += cell.GetRate()
					day1Count++
					break
				}
			}
		}
		avgDay1 := 0.0
		if day1Count > 0 {
			avgDay1 = day1Sum / float64(day1Count)
		}
		summary := fmt.Sprintf("共 %d 个同期群，平均次日留存 %.2f%%。", len(cohorts), avgDay1*100)

		headers := []string{"cohort_date", "size", "offset_days", "count", "rate"}
		rows := make([][]string, 0, len(cohorts)*4)
		for _, c := range cohorts {
			for _, cell := range c.GetCells() {
				rows = append(rows, []string{
					time.UnixMilli(c.GetCohortDate()).Format("2006-01-02"),
					fmt.Sprintf("%d", c.GetSize()),
					fmt.Sprintf("%d", cell.GetOffsetDays()),
					fmt.Sprintf("%d", cell.GetCount()),
					fmt.Sprintf("%.4f", cell.GetRate()),
				})
			}
		}
		return summary, csvFromRows(headers, rows), nil

	default:
		return "", "", fmt.Errorf("unsupported report model: %s", derefStr(sub.Model))
	}
}

// deliver 把报告摘要推送到订阅的渠道。
func (s *ReportEngine) deliver(ctx context.Context, sub *ent.ReportSubscription, summary, csv string, now time.Time) {
	channels := sub.Channels
	if len(channels) == 0 {
		channels = []string{"INTERNAL_MESSAGE"}
	}

	title := fmt.Sprintf("【定时报告】%s", derefStr(sub.Name))

	for _, channel := range channels {
		switch channel {
		case "INTERNAL_MESSAGE":
			s.sendInternalMessage(ctx, sub, title, summary, csv, now)
		case "WEBHOOK":
			// 应用维度不传：webhook 的应用过滤由其自身配置决定。
			s.webhookSender.Send(ctx, SendPayload{
				TenantID:  derefU32(sub.TenantID),
				AppID:     "",
				EventType: reportEventType,
				Data: map[string]any{
					"subscriptionId": sub.ID,
					"subscription":   derefStr(sub.Name),
					"model":          derefStr(sub.Model),
					"summary":        summary,
					"csv":            csv,
					"generatedAt":    now.Format(time.RFC3339),
				},
			})
		}
	}
}

func (s *ReportEngine) sendInternalMessage(ctx context.Context, sub *ent.ReportSubscription, title, summary, csv string, now time.Time) {
	tenantID := derefU32(sub.TenantID)

	content := summary
	if csv != "" {
		content = summary + "\n\n" + csv
	}

	msg, err := s.msgRepo.Create(ctx, &internalMessageV1.CreateInternalMessageRequest{
		Data: &internalMessageV1.InternalMessage{
			TenantId:  &tenantID,
			Title:     &title,
			Content:   &content,
			CreatedAt: timestamppb.New(now),
		},
	})
	if err != nil {
		s.log.Errorf("create report internal message failed: %s", err.Error())
		return
	}

	if sub.CreatedBy == nil {
		return
	}

	recipient := *sub.CreatedBy
	if _, err = s.recipientRepo.Create(ctx, &internalMessageV1.InternalMessageRecipient{
		MessageId:       msg.Id,
		RecipientUserId: &recipient,
		TenantId:        &tenantID,
	}); err != nil {
		s.log.Errorf("create report message recipient failed: %s", err.Error())
	}
}

func (s *ReportEngine) analyticsImpl() reportAnalytics {
	if data.UseClickHouse {
		return s.chAnalytics
	}
	return s.analytics
}

// parseReportParams 把订阅里的参数 JSON 快照解析为对应分析请求；空参数保持零值。
func parseReportParams(raw string, req proto.Message) error {
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	opts := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err := opts.Unmarshal([]byte(raw), req); err != nil {
		return fmt.Errorf("invalid report params: %w", err)
	}
	return nil
}

// ensureReportRange 参数缺时间范围时按执行间隔回看（新建订阅未配范围也能跑出报告）。
func ensureReportRange(tr *ubaV1.TimeRange, interval time.Duration, now time.Time) *ubaV1.TimeRange {
	if tr != nil && tr.GetStartMs() != 0 {
		return tr
	}
	return &ubaV1.TimeRange{
		StartMs: now.Add(-interval).UnixMilli(),
		EndMs:   now.UnixMilli(),
	}
}

// csvFromRows 把表头与行渲染为 CSV 字符串（RFC4180 引号转义）。
func csvFromRows(headers []string, rows [][]string) string {
	var sb strings.Builder
	writeRow := func(cells []string) {
		quoted := make([]string, 0, len(cells))
		for _, c := range cells {
			q := strings.ReplaceAll(c, "\"", "\"\"")
			q = strings.ReplaceAll(q, "\n", " ")
			quoted = append(quoted, "\""+q+"\"")
		}
		sb.WriteString(strings.Join(quoted, ","))
		sb.WriteString("\n")
	}
	writeRow(headers)
	for i, row := range rows {
		if i >= reportCSVMaxRows {
			break
		}
		writeRow(row)
	}
	return sb.String()
}
