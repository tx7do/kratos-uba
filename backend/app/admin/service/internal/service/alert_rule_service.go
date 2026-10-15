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

type AlertRuleService struct {
	adminV1.AlertRuleServiceHTTPServer

	log *log.Helper

	alertRuleServiceClient ubaV1.AlertRuleServiceClient
}

func NewAlertRuleService(
	ctx *bootstrap.Context,
	alertRuleServiceClient ubaV1.AlertRuleServiceClient,
) *AlertRuleService {
	return &AlertRuleService{
		log:                    ctx.NewLoggerHelper("alert-rule/service/admin-service"),
		alertRuleServiceClient: alertRuleServiceClient,
	}
}

func (s *AlertRuleService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListAlertRuleResponse, error) {
	return s.alertRuleServiceClient.List(ctx, req)
}

func (s *AlertRuleService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountAlertRuleResponse, error) {
	return s.alertRuleServiceClient.Count(ctx, req)
}

func (s *AlertRuleService) Get(ctx context.Context, req *ubaV1.GetAlertRuleRequest) (*ubaV1.AlertRule, error) {
	return s.alertRuleServiceClient.Get(ctx, req)
}

func (s *AlertRuleService) Create(ctx context.Context, req *ubaV1.CreateAlertRuleRequest) (*ubaV1.AlertRule, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.alertRuleServiceClient.Create(ctx, req)
}

func (s *AlertRuleService) Update(ctx context.Context, req *ubaV1.UpdateAlertRuleRequest) (*ubaV1.AlertRule, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.alertRuleServiceClient.Update(ctx, req)
}

func (s *AlertRuleService) Delete(ctx context.Context, req *ubaV1.DeleteAlertRuleRequest) (*emptypb.Empty, error) {
	return s.alertRuleServiceClient.Delete(ctx, req)
}
