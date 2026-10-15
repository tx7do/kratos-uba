package service

import (
	"context"

	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

type FeatureFlagService struct {
	adminV1.FeatureFlagServiceHTTPServer

	log *log.Helper

	featureFlagServiceClient ubaV1.FeatureFlagServiceClient
}

func NewFeatureFlagService(
	ctx *bootstrap.Context,
	featureFlagServiceClient ubaV1.FeatureFlagServiceClient,
) *FeatureFlagService {
	return &FeatureFlagService{
		log:                      ctx.NewLoggerHelper("feature-flag/service/admin-service"),
		featureFlagServiceClient: featureFlagServiceClient,
	}
}

func (s *FeatureFlagService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListFeatureFlagResponse, error) {
	return s.featureFlagServiceClient.List(ctx, req)
}

func (s *FeatureFlagService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountFeatureFlagResponse, error) {
	return s.featureFlagServiceClient.Count(ctx, req)
}

func (s *FeatureFlagService) Get(ctx context.Context, req *ubaV1.GetFeatureFlagRequest) (*ubaV1.FeatureFlag, error) {
	return s.featureFlagServiceClient.Get(ctx, req)
}

func (s *FeatureFlagService) Create(ctx context.Context, req *ubaV1.CreateFeatureFlagRequest) (*ubaV1.FeatureFlag, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.featureFlagServiceClient.Create(ctx, req)
}

func (s *FeatureFlagService) Update(ctx context.Context, req *ubaV1.UpdateFeatureFlagRequest) (*ubaV1.FeatureFlag, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.featureFlagServiceClient.Update(ctx, req)
}

func (s *FeatureFlagService) Delete(ctx context.Context, req *ubaV1.DeleteFeatureFlagRequest) (*emptypb.Empty, error) {
	return s.featureFlagServiceClient.Delete(ctx, req)
}

// GetFlagResult 实验结果：各变体曝光人数与目标事件转化率。
func (s *FeatureFlagService) GetFlagResult(ctx context.Context, req *ubaV1.GetFlagResultRequest) (*ubaV1.FlagResultResponse, error) {
	return s.featureFlagServiceClient.GetFlagResult(ctx, req)
}
