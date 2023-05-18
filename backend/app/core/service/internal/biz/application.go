package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-bi/gen/api/go/common/pagination"
	v1 "kratos-bi/gen/api/go/user/service/v1"
)

type ApplicationRepo interface {
	Create(ctx context.Context, req *v1.CreateApplicationRequest) (*v1.Application, error)
	Update(ctx context.Context, req *v1.UpdateApplicationRequest) (*v1.Application, error)
	Delete(ctx context.Context, req *v1.DeleteApplicationRequest) (bool, error)
	List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListApplicationResponse, error)
	Get(ctx context.Context, req *v1.GetApplicationRequest) (*v1.Application, error)
}

type ApplicationUseCase struct {
	repo ApplicationRepo
	log  *log.Helper
}

func NewApplicationUseCase(repo ApplicationRepo, logger log.Logger) *ApplicationUseCase {
	l := log.NewHelper(log.With(logger, "module", "application/usecase/core-service"))
	return &ApplicationUseCase{
		repo: repo,
		log:  l,
	}
}

func (uc *ApplicationUseCase) Get(ctx context.Context, req *v1.GetApplicationRequest) (*v1.Application, error) {
	user, err := uc.repo.Get(ctx, req)
	return user, err
}

func (uc *ApplicationUseCase) List(ctx context.Context, req *pagination.PagingRequest) (*v1.ListApplicationResponse, error) {
	return uc.repo.List(ctx, req)
}

func (uc *ApplicationUseCase) Create(ctx context.Context, req *v1.CreateApplicationRequest) (*v1.Application, error) {
	return uc.repo.Create(ctx, req)
}

func (uc *ApplicationUseCase) Update(ctx context.Context, req *v1.UpdateApplicationRequest) (*v1.Application, error) {
	return uc.repo.Update(ctx, req)
}

func (uc *ApplicationUseCase) Delete(ctx context.Context, req *v1.DeleteApplicationRequest) (bool, error) {
	return uc.repo.Delete(ctx, req)
}
