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

type ReportSubscriptionService struct {
	ubaV1.UnimplementedReportSubscriptionServiceServer

	log *log.Helper

	reportRepo *data.ReportSubscriptionRepo
}

func NewReportSubscriptionService(
	ctx *bootstrap.Context,
	reportRepo *data.ReportSubscriptionRepo,
) *ReportSubscriptionService {
	return &ReportSubscriptionService{
		log:        ctx.NewLoggerHelper("report-subscription/service/core-service"),
		reportRepo: reportRepo,
	}
}

func (s *ReportSubscriptionService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListReportSubscriptionResponse, error) {
	return s.reportRepo.List(ctx, req)
}

func (s *ReportSubscriptionService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountReportSubscriptionResponse, error) {
	return s.reportRepo.Count(ctx, req)
}

func (s *ReportSubscriptionService) Get(ctx context.Context, req *ubaV1.GetReportSubscriptionRequest) (*ubaV1.ReportSubscription, error) {
	return s.reportRepo.Get(ctx, req)
}

func (s *ReportSubscriptionService) Create(ctx context.Context, req *ubaV1.CreateReportSubscriptionRequest) (*ubaV1.ReportSubscription, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.reportRepo.Create(ctx, req)
}

func (s *ReportSubscriptionService) Update(ctx context.Context, req *ubaV1.UpdateReportSubscriptionRequest) (*ubaV1.ReportSubscription, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.reportRepo.Update(ctx, req)
}

func (s *ReportSubscriptionService) Delete(ctx context.Context, req *ubaV1.DeleteReportSubscriptionRequest) (*emptypb.Empty, error) {
	if err := s.reportRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
