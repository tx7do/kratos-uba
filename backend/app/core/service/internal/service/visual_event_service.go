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

// VisualEventService 可视化埋点规则：页面 + 元素 XPath → 事件名的配置 CRUD，
// 以及为 collector /uba/v1/visual-events 下发端点提供启用的规则列表。
type VisualEventService struct {
	ubaV1.UnimplementedVisualEventServiceServer

	log *log.Helper

	visualEventRepo *data.VisualEventRepo
}

func NewVisualEventService(
	ctx *bootstrap.Context,
	visualEventRepo *data.VisualEventRepo,
) *VisualEventService {
	return &VisualEventService{
		log:             ctx.NewLoggerHelper("visual-event/service/core-service"),
		visualEventRepo: visualEventRepo,
	}
}

func (s *VisualEventService) List(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.ListVisualEventResponse, error) {
	return s.visualEventRepo.List(ctx, req)
}

func (s *VisualEventService) Count(ctx context.Context, req *paginationV1.PagingRequest) (*ubaV1.CountVisualEventResponse, error) {
	return s.visualEventRepo.Count(ctx, req)
}

func (s *VisualEventService) Get(ctx context.Context, req *ubaV1.GetVisualEventRequest) (*ubaV1.VisualEvent, error) {
	return s.visualEventRepo.Get(ctx, req)
}

func (s *VisualEventService) Create(ctx context.Context, req *ubaV1.CreateVisualEventRequest) (*ubaV1.VisualEvent, error) {
	if req == nil || req.Data == nil {
		return nil, ubaV1.ErrorVisualEventBadRequest("invalid parameter")
	}
	if req.Data.GetEventName() == "" {
		return nil, ubaV1.ErrorVisualEventBadRequest("event_name is required")
	}
	if req.Data.GetElementXpath() == "" {
		return nil, ubaV1.ErrorVisualEventBadRequest("element_xpath is required")
	}

	return s.visualEventRepo.Create(ctx, req)
}

func (s *VisualEventService) Update(ctx context.Context, req *ubaV1.UpdateVisualEventRequest) (*ubaV1.VisualEvent, error) {
	if req == nil || req.Data == nil {
		return nil, ubaV1.ErrorVisualEventBadRequest("invalid parameter")
	}

	return s.visualEventRepo.Update(ctx, req)
}

func (s *VisualEventService) Delete(ctx context.Context, req *ubaV1.DeleteVisualEventRequest) (*emptypb.Empty, error) {
	if err := s.visualEventRepo.Delete(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// ListForApp 为应用下发启用的规则。collector 已做应用鉴权并带权威租户 ID。
func (s *VisualEventService) ListForApp(ctx context.Context, req *ubaV1.ListVisualEventsForAppRequest) (*ubaV1.ListVisualEventsForAppResponse, error) {
	if req == nil || req.GetTenantId() == 0 {
		return nil, ubaV1.ErrorVisualEventBadRequest("tenant_id is required")
	}

	entities, err := s.visualEventRepo.ListEnabledByTenantApp(ctx, req.GetTenantId(), req.GetAppId())
	if err != nil {
		return nil, err
	}

	items := make([]*ubaV1.VisualEvent, 0, len(entities))
	for i := range entities {
		entity := entities[i]
		enabled := entity.Enabled
		appID := derefStr(entity.AppID)
		id := entity.ID
		items = append(items, &ubaV1.VisualEvent{
			Id:           id,
			TenantId:     entity.TenantID,
			AppId:        &appID,
			EventName:    entity.EventName,
			DisplayName:  entity.DisplayName,
			PageUrl:      entity.PageURL,
			ElementXpath: &entity.ElementXpath,
			MatchType:    matchTypeToDTO(entity.MatchType),
			Properties:   entity.Properties,
			Enabled:      &enabled,
		})
	}

	return &ubaV1.ListVisualEventsForAppResponse{Items: items}, nil
}

// matchTypeToDTO 存储字符串转枚举（未知值回退精确匹配）。
func matchTypeToDTO(s string) *ubaV1.VisualMatchType {
	if v, ok := ubaV1.VisualMatchType_value[s]; ok {
		mt := ubaV1.VisualMatchType(v)
		return &mt
	}
	mt := ubaV1.VisualMatchType_XPATH_EXACT
	return &mt
}
