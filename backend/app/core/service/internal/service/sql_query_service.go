package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	dorisData "go-wind-uba/app/core/service/internal/data/doris"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

// SqlQueryService 自助 SQL 查询：对 OLAP 分析库的受控只读通道。
// 安全边界（见 doris.ValidateAdhocQuery / ExecuteAdhocQuery）：
// 仅允许单条 SELECT/SHOW/DESCRIBE/EXPLAIN、拒绝注释与多语句、
// 强制行数上限（默认 200 / 硬顶 1000）与服务端 30s 超时。
type SqlQueryService struct {
	ubaV1.UnimplementedSqlQueryServiceServer

	log       *log.Helper
	analytics *dorisData.AnalyticsRepo
}

func NewSqlQueryService(
	ctx *bootstrap.Context,
	analytics *dorisData.AnalyticsRepo,
) *SqlQueryService {
	return &SqlQueryService{
		log:       ctx.NewLoggerHelper("sql-query/service/core-service"),
		analytics: analytics,
	}
}

// ExecuteQuery 执行只读查询。
func (s *SqlQueryService) ExecuteQuery(ctx context.Context, req *ubaV1.ExecuteSqlQueryRequest) (*ubaV1.ExecuteSqlQueryResponse, error) {
	if req == nil || req.GetQuery() == "" {
		return nil, ubaV1.ErrorBadRequest("query is required")
	}

	resp, err := s.analytics.ExecuteAdhocQuery(ctx, req.GetQuery(), req.GetMaxRows())
	if err != nil {
		return nil, err
	}

	s.log.Infof("adhoc query executed: rows=%d truncated=%v duration=%dms query=%q",
		resp.GetRowCount(), resp.GetTruncated(), resp.GetDurationMs(), req.GetQuery())

	return resp, nil
}
