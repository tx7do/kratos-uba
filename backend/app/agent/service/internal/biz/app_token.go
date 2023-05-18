package biz

import (
	"context"

	v1 "kratos-bi/gen/api/go/user/service/v1"
)

type ApplicationTokenRepo interface {
	GetToken(ctx context.Context, appId uint32) string
	GenerateToken(ctx context.Context, app *v1.Application) (string, error)
	RemoveToken(ctx context.Context, appId uint32) error
}

type ApplicationTokenUseCase struct {
	repo ApplicationTokenRepo
}

func NewApplicationAuthUseCase(repo ApplicationTokenRepo) *ApplicationTokenUseCase {
	return &ApplicationTokenUseCase{
		repo: repo,
	}
}

func (uc *ApplicationTokenUseCase) GenerateToken(ctx context.Context, app *v1.Application) (string, error) {
	return uc.repo.GenerateToken(ctx, app)
}

func (uc *ApplicationTokenUseCase) GetToken(ctx context.Context, appId uint32) string {
	return uc.repo.GetToken(ctx, appId)
}

func (uc *ApplicationTokenUseCase) RemoveToken(ctx context.Context, appId uint32) error {
	return uc.repo.RemoveToken(ctx, appId)
}
