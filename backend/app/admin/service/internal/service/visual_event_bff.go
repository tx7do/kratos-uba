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

type VisualEventService struct {
	adminV1.VisualEventServiceHTTPServer

	log *log.Helper

	visualEventServiceClient ubaV1.VisualEventServiceClient
}

func NewVisualEventService(
	ctx *bootstrap.Context,
	visualEventServiceClient ubaV1.VisualEventServiceClient,
) *VisualEventService {
	return &VisualEventService{
		log:                      ctx.NewLoggerHelper("visual-event/service/admin-service"),
		visualEventServiceClient: visualEventServiceClient,
	}
}

func (s *VisualEventService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListVisualEventResponse, error) {
	return s.visualEventServiceClient.List(ctx, req)
}

func (s *VisualEventService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountVisualEventResponse, error) {
	return s.visualEventServiceClient.Count(ctx, req)
}

func (s *VisualEventService) Get(ctx context.Context, req *ubaV1.GetVisualEventRequest) (*ubaV1.VisualEvent, error) {
	return s.visualEventServiceClient.Get(ctx, req)
}

func (s *VisualEventService) Create(ctx context.Context, req *ubaV1.CreateVisualEventRequest) (*ubaV1.VisualEvent, error) {
	if req == nil || req.Data == nil {
		return nil, ubaV1.ErrorVisualEventBadRequest("invalid parameter")
	}

	return s.visualEventServiceClient.Create(ctx, req)
}

func (s *VisualEventService) Update(ctx context.Context, req *ubaV1.UpdateVisualEventRequest) (*ubaV1.VisualEvent, error) {
	if req == nil || req.Data == nil {
		return nil, ubaV1.ErrorVisualEventBadRequest("invalid parameter")
	}

	return s.visualEventServiceClient.Update(ctx, req)
}

func (s *VisualEventService) Delete(ctx context.Context, req *ubaV1.DeleteVisualEventRequest) (*emptypb.Empty, error) {
	return s.visualEventServiceClient.Delete(ctx, req)
}

// ListForApp 为应用下发启用规则（collector 转发；管理端调试用）。
func (s *VisualEventService) ListForApp(ctx context.Context, req *ubaV1.ListVisualEventsForAppRequest) (*ubaV1.ListVisualEventsForAppResponse, error) {
	return s.visualEventServiceClient.ListForApp(ctx, req)
}
