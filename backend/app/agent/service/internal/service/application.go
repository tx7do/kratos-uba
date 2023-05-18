package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "kratos-bi/gen/api/go/agent/service/v1"
	"kratos-bi/gen/api/go/common/pagination"
	userV1 "kratos-bi/gen/api/go/user/service/v1"
)

type ApplicationService struct {
	v1.ApplicationServiceHTTPServer

	tagClient userV1.ApplicationServiceClient
	log       *log.Helper
}

func NewApplicationService(logger log.Logger, tagClient userV1.ApplicationServiceClient) *ApplicationService {
	l := log.NewHelper(log.With(logger, "module", "app/service/agent-service"))
	return &ApplicationService{
		log:       l,
		tagClient: tagClient,
	}
}

// ListApplication 获取标签列表
func (s *ApplicationService) ListApplication(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListApplicationResponse, error) {
	return s.tagClient.ListApplication(ctx, req)
}

// GetApplication 获取标签数据
func (s *ApplicationService) GetApplication(ctx context.Context, req *userV1.GetApplicationRequest) (*userV1.Application, error) {
	return s.tagClient.GetApplication(ctx, req)
}

// CreateApplication 创建标签
func (s *ApplicationService) CreateApplication(ctx context.Context, req *userV1.CreateApplicationRequest) (*userV1.Application, error) {
	return s.tagClient.CreateApplication(ctx, req)
}

// UpdateApplication 更新标签
func (s *ApplicationService) UpdateApplication(ctx context.Context, req *userV1.UpdateApplicationRequest) (*userV1.Application, error) {
	return s.tagClient.UpdateApplication(ctx, req)
}

// DeleteApplication 删除标签
func (s *ApplicationService) DeleteApplication(ctx context.Context, req *userV1.DeleteApplicationRequest) (*emptypb.Empty, error) {
	return s.tagClient.DeleteApplication(ctx, req)
}
