package service

import (
	"context"

	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	paginationV1 "github.com/tx7do/go-crud/api/gen/go/pagination/v1"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	notificationChannelV1 "go-wind-uba/api/gen/go/notification_channel/service/v1"
	notificationV1 "go-wind-uba/api/gen/go/notification/service/v1"
)

// ── NotificationService（投递台账只读 + SendDirect 不开放 HTTP） ──

type NotificationService struct {
	adminV1.NotificationServiceHTTPServer

	log *log.Helper

	notificationServiceClient notificationV1.NotificationServiceClient
}

func NewNotificationService(
	ctx *bootstrap.Context,
	notificationServiceClient notificationV1.NotificationServiceClient,
) *NotificationService {
	return &NotificationService{
		log:                       ctx.NewLoggerHelper("notification/service/admin-service"),
		notificationServiceClient: notificationServiceClient,
	}
}

func (s *NotificationService) ListNotificationDelivery(ctx context.Context, req *paginationV1.PagingRequest) (*notificationV1.ListNotificationDeliveryResponse, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification delivery ledger"); err != nil {
		return nil, err
	}
	return s.notificationServiceClient.ListDelivery(ctx, req)
}

func (s *NotificationService) GetNotificationDelivery(ctx context.Context, req *notificationV1.GetNotificationDeliveryRequest) (*notificationV1.NotificationDelivery, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification delivery ledger"); err != nil {
		return nil, err
	}
	return s.notificationServiceClient.GetDelivery(ctx, req)
}

// ── NotificationRuleService ──

type NotificationRuleService struct {
	adminV1.NotificationRuleServiceHTTPServer

	log *log.Helper

	notificationRuleServiceClient notificationV1.NotificationRuleServiceClient
}

func NewNotificationRuleService(
	ctx *bootstrap.Context,
	notificationRuleServiceClient notificationV1.NotificationRuleServiceClient,
) *NotificationRuleService {
	return &NotificationRuleService{
		log:                           ctx.NewLoggerHelper("notification-rule/service/admin-service"),
		notificationRuleServiceClient: notificationRuleServiceClient,
	}
}

func (s *NotificationRuleService) ListNotificationRule(ctx context.Context, req *paginationV1.PagingRequest) (*notificationV1.ListNotificationRuleResponse, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification routing rules"); err != nil {
		return nil, err
	}
	return s.notificationRuleServiceClient.List(ctx, req)
}

func (s *NotificationRuleService) GetNotificationRule(ctx context.Context, req *notificationV1.GetNotificationRuleRequest) (*notificationV1.NotificationRule, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification routing rules"); err != nil {
		return nil, err
	}
	return s.notificationRuleServiceClient.Get(ctx, req)
}

func (s *NotificationRuleService) CreateNotificationRule(ctx context.Context, req *notificationV1.CreateNotificationRuleRequest) (*notificationV1.NotificationRule, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification routing rules"); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, notificationV1.ErrorNotificationBadRequest("invalid parameter")
	}

	return s.notificationRuleServiceClient.Create(ctx, req)
}

func (s *NotificationRuleService) UpdateNotificationRule(ctx context.Context, req *notificationV1.UpdateNotificationRuleRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification routing rules"); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, notificationV1.ErrorNotificationBadRequest("invalid parameter")
	}

	return s.notificationRuleServiceClient.Update(ctx, req)
}

func (s *NotificationRuleService) DeleteNotificationRule(ctx context.Context, req *notificationV1.DeleteNotificationRuleRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification routing rules"); err != nil {
		return nil, err
	}
	return s.notificationRuleServiceClient.Delete(ctx, req)
}

func (s *NotificationRuleService) TestDispatchNotification(ctx context.Context, req *notificationV1.TestDispatchNotificationRequest) (*notificationV1.TestDispatchNotificationResponse, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification routing rules"); err != nil {
		return nil, err
	}
	return s.notificationRuleServiceClient.TestDispatch(ctx, req)
}

// ── NotificationTemplateService ──

type NotificationTemplateService struct {
	adminV1.NotificationTemplateServiceHTTPServer

	log *log.Helper

	notificationTemplateServiceClient notificationV1.NotificationTemplateServiceClient
}

func NewNotificationTemplateService(
	ctx *bootstrap.Context,
	notificationTemplateServiceClient notificationV1.NotificationTemplateServiceClient,
) *NotificationTemplateService {
	return &NotificationTemplateService{
		log:                               ctx.NewLoggerHelper("notification-template/service/admin-service"),
		notificationTemplateServiceClient: notificationTemplateServiceClient,
	}
}

func (s *NotificationTemplateService) ListNotificationTemplate(ctx context.Context, req *paginationV1.PagingRequest) (*notificationV1.ListNotificationTemplateResponse, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification templates"); err != nil {
		return nil, err
	}
	return s.notificationTemplateServiceClient.ListNotificationTemplate(ctx, req)
}

func (s *NotificationTemplateService) GetNotificationTemplate(ctx context.Context, req *notificationV1.GetNotificationTemplateRequest) (*notificationV1.NotificationTemplate, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification templates"); err != nil {
		return nil, err
	}
	return s.notificationTemplateServiceClient.GetNotificationTemplate(ctx, req)
}

func (s *NotificationTemplateService) CreateNotificationTemplate(ctx context.Context, req *notificationV1.CreateNotificationTemplateRequest) (*notificationV1.NotificationTemplate, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification templates"); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, notificationV1.ErrorNotificationBadRequest("invalid parameter")
	}

	return s.notificationTemplateServiceClient.CreateNotificationTemplate(ctx, req)
}

func (s *NotificationTemplateService) UpdateNotificationTemplate(ctx context.Context, req *notificationV1.UpdateNotificationTemplateRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification templates"); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, notificationV1.ErrorNotificationBadRequest("invalid parameter")
	}

	return s.notificationTemplateServiceClient.UpdateNotificationTemplate(ctx, req)
}

func (s *NotificationTemplateService) DeleteNotificationTemplate(ctx context.Context, req *notificationV1.DeleteNotificationTemplateRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification templates"); err != nil {
		return nil, err
	}
	return s.notificationTemplateServiceClient.DeleteNotificationTemplate(ctx, req)
}

func (s *NotificationTemplateService) RenderNotificationTemplate(ctx context.Context, req *notificationV1.RenderNotificationTemplateRequest) (*notificationV1.RenderNotificationTemplateResponse, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification templates"); err != nil {
		return nil, err
	}
	return s.notificationTemplateServiceClient.RenderNotificationTemplate(ctx, req)
}

// ── NotificationChannelService ──

type NotificationChannelService struct {
	adminV1.NotificationChannelServiceHTTPServer

	log *log.Helper

	notificationChannelServiceClient notificationChannelV1.NotificationChannelServiceClient
}

func NewNotificationChannelService(
	ctx *bootstrap.Context,
	notificationChannelServiceClient notificationChannelV1.NotificationChannelServiceClient,
) *NotificationChannelService {
	return &NotificationChannelService{
		log:                              ctx.NewLoggerHelper("notification-channel/service/admin-service"),
		notificationChannelServiceClient: notificationChannelServiceClient,
	}
}

func (s *NotificationChannelService) ListNotificationChannel(ctx context.Context, req *paginationV1.PagingRequest) (*notificationChannelV1.ListNotificationChannelResponse, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	return s.notificationChannelServiceClient.List(ctx, req)
}

func (s *NotificationChannelService) GetNotificationChannel(ctx context.Context, req *notificationChannelV1.GetNotificationChannelRequest) (*notificationChannelV1.NotificationChannel, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	return s.notificationChannelServiceClient.Get(ctx, req)
}

func (s *NotificationChannelService) CreateNotificationChannel(ctx context.Context, req *notificationChannelV1.CreateNotificationChannelRequest) (*notificationChannelV1.NotificationChannel, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, notificationV1.ErrorNotificationBadRequest("invalid parameter")
	}

	return s.notificationChannelServiceClient.Create(ctx, req)
}

func (s *NotificationChannelService) UpdateNotificationChannel(ctx context.Context, req *notificationChannelV1.UpdateNotificationChannelRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, notificationV1.ErrorNotificationBadRequest("invalid parameter")
	}

	return s.notificationChannelServiceClient.Update(ctx, req)
}

func (s *NotificationChannelService) DeleteNotificationChannel(ctx context.Context, req *notificationChannelV1.DeleteNotificationChannelRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	return s.notificationChannelServiceClient.Delete(ctx, req)
}

func (s *NotificationChannelService) SendTestEmail(ctx context.Context, req *notificationChannelV1.SendTestEmailRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	return s.notificationChannelServiceClient.SendTestEmail(ctx, req)
}

func (s *NotificationChannelService) SendTest(ctx context.Context, req *notificationChannelV1.SendTestRequest) (*emptypb.Empty, error) {
	if err := requirePlatformAdmin(ctx, s.log, "notification channels"); err != nil {
		return nil, err
	}
	return s.notificationChannelServiceClient.SendTest(ctx, req)
}

// ── NotificationPreferenceService（自助：服务端从操作人钉定 user） ──

type NotificationPreferenceService struct {
	adminV1.NotificationPreferenceServiceHTTPServer

	log *log.Helper

	notificationPreferenceServiceClient notificationV1.NotificationPreferenceServiceClient
}

func NewNotificationPreferenceService(
	ctx *bootstrap.Context,
	notificationPreferenceServiceClient notificationV1.NotificationPreferenceServiceClient,
) *NotificationPreferenceService {
	return &NotificationPreferenceService{
		log:                                  ctx.NewLoggerHelper("notification-preference/service/admin-service"),
		notificationPreferenceServiceClient:  notificationPreferenceServiceClient,
	}
}

func (s *NotificationPreferenceService) GetMyNotificationPreference(ctx context.Context, req *emptypb.Empty) (*notificationV1.NotificationPreference, error) {
	return s.notificationPreferenceServiceClient.GetMyNotificationPreference(ctx, req)
}

func (s *NotificationPreferenceService) UpdateMyNotificationPreference(ctx context.Context, req *notificationV1.UpdateNotificationPreferenceRequest) (*notificationV1.NotificationPreference, error) {
	return s.notificationPreferenceServiceClient.UpdateMyNotificationPreference(ctx, req)
}

func (s *NotificationPreferenceService) ListMyNotificationCategories(ctx context.Context, req *emptypb.Empty) (*notificationV1.ListMyNotificationCategoriesResponse, error) {
	return s.notificationPreferenceServiceClient.ListMyNotificationCategories(ctx, req)
}
