package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	aiV1 "go-wind-uba/api/gen/go/ai/service/v1"
	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"
)

// AiUsageLogService AI 用量流水查询（BFF 直通转发；
// 租户口径过滤在 core 侧执行：租户用户按本租户，平台用户见全量）。
type AiUsageLogService struct {
	adminV1.AiUsageLogServiceHTTPServer

	aiUsageLogServiceClient aiV1.AiUsageLogServiceClient
	log                     *log.Helper
}

func NewAiUsageLogService(ctx *bootstrap.Context, aiUsageLogServiceClient aiV1.AiUsageLogServiceClient) *AiUsageLogService {
	return &AiUsageLogService{
		log:                     ctx.NewLoggerHelper("ai_usage_log/service/admin-service"),
		aiUsageLogServiceClient: aiUsageLogServiceClient,
	}
}

func (s *AiUsageLogService) List(ctx context.Context, req *paginationV1.PagingRequest) (*aiV1.ListAiUsageLogResponse, error) {
	return s.aiUsageLogServiceClient.List(ctx, req)
}

func (s *AiUsageLogService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*aiV1.CountAiUsageLogResponse, error) {
	return s.aiUsageLogServiceClient.Count(ctx, req)
}

func (s *AiUsageLogService) GetUsageSummary(ctx context.Context, req *emptypb.Empty) (*aiV1.UsageSummaryResponse, error) {
	return s.aiUsageLogServiceClient.GetUsageSummary(ctx, req)
}

// GetTenantQuota 查询租户月度 token 配额（租户用户在 core 侧被强制看本租户）。
func (s *AiUsageLogService) GetTenantQuota(ctx context.Context, req *aiV1.GetTenantQuotaRequest) (*aiV1.TenantQuotaResponse, error) {
	return s.aiUsageLogServiceClient.GetTenantQuota(ctx, req)
}

// SetTenantQuota 设置租户月度 token 配额（平台管理员专属；limit=0 取消）。
func (s *AiUsageLogService) SetTenantQuota(ctx context.Context, req *aiV1.SetTenantQuotaRequest) (*emptypb.Empty, error) {
	return s.aiUsageLogServiceClient.SetTenantQuota(ctx, req)
}
