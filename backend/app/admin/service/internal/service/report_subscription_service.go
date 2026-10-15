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

type ReportSubscriptionService struct {
	adminV1.ReportSubscriptionServiceHTTPServer

	log *log.Helper

	reportSubscriptionServiceClient ubaV1.ReportSubscriptionServiceClient
}

func NewReportSubscriptionService(
	ctx *bootstrap.Context,
	reportSubscriptionServiceClient ubaV1.ReportSubscriptionServiceClient,
) *ReportSubscriptionService {
	return &ReportSubscriptionService{
		log:                             ctx.NewLoggerHelper("report-subscription/service/admin-service"),
		reportSubscriptionServiceClient: reportSubscriptionServiceClient,
	}
}

func (s *ReportSubscriptionService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListReportSubscriptionResponse, error) {
	return s.reportSubscriptionServiceClient.List(ctx, req)
}

func (s *ReportSubscriptionService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountReportSubscriptionResponse, error) {
	return s.reportSubscriptionServiceClient.Count(ctx, req)
}

func (s *ReportSubscriptionService) Get(ctx context.Context, req *ubaV1.GetReportSubscriptionRequest) (*ubaV1.ReportSubscription, error) {
	return s.reportSubscriptionServiceClient.Get(ctx, req)
}

func (s *ReportSubscriptionService) Create(ctx context.Context, req *ubaV1.CreateReportSubscriptionRequest) (*ubaV1.ReportSubscription, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.reportSubscriptionServiceClient.Create(ctx, req)
}

func (s *ReportSubscriptionService) Update(ctx context.Context, req *ubaV1.UpdateReportSubscriptionRequest) (*ubaV1.ReportSubscription, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.reportSubscriptionServiceClient.Update(ctx, req)
}

func (s *ReportSubscriptionService) Delete(ctx context.Context, req *ubaV1.DeleteReportSubscriptionRequest) (*emptypb.Empty, error) {
	return s.reportSubscriptionServiceClient.Delete(ctx, req)
}
