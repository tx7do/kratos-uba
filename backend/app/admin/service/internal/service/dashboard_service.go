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

type DashboardService struct {
	adminV1.DashboardServiceHTTPServer

	log *log.Helper

	dashboardServiceClient ubaV1.DashboardServiceClient
}

func NewDashboardService(
	ctx *bootstrap.Context,
	dashboardServiceClient ubaV1.DashboardServiceClient,
) *DashboardService {
	return &DashboardService{
		log:                    ctx.NewLoggerHelper("dashboard/service/admin-service"),
		dashboardServiceClient: dashboardServiceClient,
	}
}

func (s *DashboardService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListDashboardResponse, error) {
	return s.dashboardServiceClient.List(ctx, req)
}

func (s *DashboardService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountDashboardResponse, error) {
	return s.dashboardServiceClient.Count(ctx, req)
}

func (s *DashboardService) Get(ctx context.Context, req *ubaV1.GetDashboardRequest) (*ubaV1.Dashboard, error) {
	return s.dashboardServiceClient.Get(ctx, req)
}

func (s *DashboardService) Create(ctx context.Context, req *ubaV1.CreateDashboardRequest) (*ubaV1.Dashboard, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.dashboardServiceClient.Create(ctx, req)
}

func (s *DashboardService) Update(ctx context.Context, req *ubaV1.UpdateDashboardRequest) (*ubaV1.Dashboard, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.dashboardServiceClient.Update(ctx, req)
}

func (s *DashboardService) Delete(ctx context.Context, req *ubaV1.DeleteDashboardRequest) (*emptypb.Empty, error) {
	return s.dashboardServiceClient.Delete(ctx, req)
}
