package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-uba/app/core/service/internal/biz"

	"kratos-uba/gen/api/go/common/pagination"
	v1 "kratos-uba/gen/api/go/user/service/v1"
)

type ApplicationService struct {
	v1.UnimplementedApplicationServiceServer

	uc  *biz.ApplicationUseCase
	log *log.Helper
}

func NewApplicationService(logger log.Logger, uc *biz.ApplicationUseCase) *ApplicationService {
	l := log.NewHelper(log.With(logger, "module", "app/service/core-service"))
	return &ApplicationService{
		log: l,
		uc:  uc,
	}
}

func (s *ApplicationService) ListApplication(ctx context.Context, req *pagination.PagingRequest) (*v1.ListApplicationResponse, error) {
	resp, err := s.uc.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ApplicationService) GetApplication(ctx context.Context, req *v1.GetApplicationRequest) (*v1.Application, error) {
	resp, err := s.uc.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ApplicationService) CreateApplication(ctx context.Context, req *v1.CreateApplicationRequest) (*v1.Application, error) {
	resp, err := s.uc.Create(ctx, req)
	if err != nil {
		// s.log.Info(err)
		return nil, err
	}
	return resp, nil
}

func (s *ApplicationService) UpdateApplication(ctx context.Context, req *v1.UpdateApplicationRequest) (*v1.Application, error) {
	resp, err := s.uc.Update(ctx, req)
	if err != nil {
		// s.log.Info(err)
		return nil, err
	}
	return resp, nil
}

func (s *ApplicationService) DeleteApplication(ctx context.Context, req *v1.DeleteApplicationRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
