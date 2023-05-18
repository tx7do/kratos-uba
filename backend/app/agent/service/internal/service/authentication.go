package service

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-bi/app/agent/service/internal/biz"

	v1 "kratos-bi/gen/api/go/agent/service/v1"
	userV1 "kratos-bi/gen/api/go/user/service/v1"
)

type AuthenticationService struct {
	v1.AuthenticationServiceHTTPServer

	uc   userV1.ApplicationServiceClient
	utuc *biz.ApplicationTokenUseCase

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, uc userV1.ApplicationServiceClient, utuc *biz.ApplicationTokenUseCase) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "auth/service/agent-service"))
	return &AuthenticationService{
		log:  l,
		uc:   uc,
		utuc: utuc,
	}
}

func (s *AuthenticationService) RefreshAccessToken(ctx context.Context, req *v1.RefreshAccessTokenRequest) (*v1.RefreshAccessTokenResponse, error) {
	app, err := s.uc.GetApplicationByAppId(ctx, &userV1.GetApplicationByAppIdRequest{
		AppId: req.GetAppId(),
	})
	if err != nil {
		return nil, err
	}

	if app.GetAppKey() != req.GetAppKey() {
		return nil, errors.New("error app key")
	}

	token, err := s.utuc.GenerateToken(ctx, app)
	if err != nil {
		return nil, err
	}

	return &v1.RefreshAccessTokenResponse{
		RefreshToken: token,
	}, nil
}
