package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"

	"go-wind-uba/app/core/service/internal/data"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

type DashboardService struct {
	ubaV1.UnimplementedDashboardServiceServer

	log *log.Helper

	dashboardRepo *data.DashboardRepo
}

func NewDashboardService(
	ctx *bootstrap.Context,
	dashboardRepo *data.DashboardRepo,
) *DashboardService {
	return &DashboardService{
		log:           ctx.NewLoggerHelper("dashboard/service/core-service"),
		dashboardRepo: dashboardRepo,
	}
}

func (s *DashboardService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListDashboardResponse, error) {
	return s.dashboardRepo.List(ctx, req)
}

func (s *DashboardService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountDashboardResponse, error) {
	return s.dashboardRepo.Count(ctx, req)
}

func (s *DashboardService) Get(ctx context.Context, req *ubaV1.GetDashboardRequest) (*ubaV1.Dashboard, error) {
	return s.dashboardRepo.Get(ctx, req)
}

func (s *DashboardService) Create(ctx context.Context, req *ubaV1.CreateDashboardRequest) (*ubaV1.Dashboard, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.dashboardRepo.Create(ctx, req)
}

func (s *DashboardService) Update(ctx context.Context, req *ubaV1.UpdateDashboardRequest) (*ubaV1.Dashboard, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.dashboardRepo.Update(ctx, req)
}

func (s *DashboardService) Delete(ctx context.Context, req *ubaV1.DeleteDashboardRequest) (*emptypb.Empty, error) {
	if err := s.dashboardRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
