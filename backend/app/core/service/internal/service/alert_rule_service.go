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

type AlertRuleService struct {
	ubaV1.UnimplementedAlertRuleServiceServer

	log *log.Helper

	alertRepo *data.AlertRuleRepo
}

func NewAlertRuleService(
	ctx *bootstrap.Context,
	alertRepo *data.AlertRuleRepo,
) *AlertRuleService {
	return &AlertRuleService{
		log:       ctx.NewLoggerHelper("alert-rule/service/core-service"),
		alertRepo: alertRepo,
	}
}

func (s *AlertRuleService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListAlertRuleResponse, error) {
	return s.alertRepo.List(ctx, req)
}

func (s *AlertRuleService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountAlertRuleResponse, error) {
	return s.alertRepo.Count(ctx, req)
}

func (s *AlertRuleService) Get(ctx context.Context, req *ubaV1.GetAlertRuleRequest) (*ubaV1.AlertRule, error) {
	return s.alertRepo.Get(ctx, req)
}

func (s *AlertRuleService) Create(ctx context.Context, req *ubaV1.CreateAlertRuleRequest) (*ubaV1.AlertRule, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.alertRepo.Create(ctx, req)
}

func (s *AlertRuleService) Update(ctx context.Context, req *ubaV1.UpdateAlertRuleRequest) (*ubaV1.AlertRule, error) {
	if req.Data == nil {
		return nil, ubaV1.ErrorBadRequest("invalid parameter")
	}

	return s.alertRepo.Update(ctx, req)
}

func (s *AlertRuleService) Delete(ctx context.Context, req *ubaV1.DeleteAlertRuleRequest) (*emptypb.Empty, error) {
	if err := s.alertRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
