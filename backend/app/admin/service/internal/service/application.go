package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	adminV1 "kratos-uba/api/gen/go/admin/service/v1"
	userV1 "kratos-uba/api/gen/go/user/service/v1"
)

type ApplicationService struct {
	adminV1.ApplicationServiceHTTPServer

	appClient userV1.ApplicationServiceClient
	log       *log.Helper
}

func NewApplicationService(logger log.Logger, appClient userV1.ApplicationServiceClient) *ApplicationService {
	l := log.NewHelper(log.With(logger, "module", "app/service/admin-service"))
	return &ApplicationService{
		log:       l,
		appClient: appClient,
	}
}

// ListApplication 获取应用列表
func (s *ApplicationService) ListApplication(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListApplicationResponse, error) {
	return s.appClient.ListApplication(ctx, req)
}

// GetApplication 获取应用数据
func (s *ApplicationService) GetApplication(ctx context.Context, req *userV1.GetApplicationRequest) (*userV1.Application, error) {
	return s.appClient.GetApplication(ctx, req)
}

// CreateApplication 创建应用
func (s *ApplicationService) CreateApplication(ctx context.Context, req *userV1.CreateApplicationRequest) (*userV1.Application, error) {
	return s.appClient.CreateApplication(ctx, req)
}

// UpdateApplication 更新应用
func (s *ApplicationService) UpdateApplication(ctx context.Context, req *userV1.UpdateApplicationRequest) (*userV1.Application, error) {
	return s.appClient.UpdateApplication(ctx, req)
}

// DeleteApplication 删除应用
func (s *ApplicationService) DeleteApplication(ctx context.Context, req *userV1.DeleteApplicationRequest) (*emptypb.Empty, error) {
	return s.appClient.DeleteApplication(ctx, req)
}
