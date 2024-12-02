package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-uba/app/admin/service/internal/data"

	adminV1 "kratos-uba/api/gen/go/admin/service/v1"
	userV1 "kratos-uba/api/gen/go/user/service/v1"

	"kratos-uba/pkg/middleware/auth"
)

type AuthenticationService struct {
	adminV1.AuthenticationServiceHTTPServer

	uc   userV1.UserServiceClient
	utuc *data.UserTokenRepo

	log *log.Helper
}

func NewAuthenticationService(logger log.Logger, uc userV1.UserServiceClient, utuc *data.UserTokenRepo) *AuthenticationService {
	l := log.NewHelper(log.With(logger, "module", "auth/service/admin-service"))
	return &AuthenticationService{
		log:  l,
		uc:   uc,
		utuc: utuc,
	}
}

// Login 登陆
func (s *AuthenticationService) Login(ctx context.Context, req *adminV1.LoginRequest) (*adminV1.LoginResponse, error) {
	if _, err := s.uc.VerifyPassword(ctx, &userV1.VerifyPasswordRequest{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
	}); err != nil {
		return &adminV1.LoginResponse{}, err
	}

	user, err := s.uc.GetUserByUserName(ctx, &userV1.GetUserByUserNameRequest{UserName: req.GetUserName()})
	if err != nil {
		return &adminV1.LoginResponse{}, err
	}

	token, err := s.utuc.GenerateToken(ctx, user)
	if err != nil {
		return &adminV1.LoginResponse{}, err
	}

	return &adminV1.LoginResponse{
		Token:    token,
		Id:       user.GetId(),
		UserName: user.GetUserName(),
	}, nil
}

// Logout 登出
func (s *AuthenticationService) Logout(ctx context.Context, req *adminV1.LogoutRequest) (*emptypb.Empty, error) {
	err := s.utuc.RemoveToken(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *AuthenticationService) GetMe(ctx context.Context, req *adminV1.GetMeRequest) (*userV1.User, error) {
	userInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, adminV1.ErrorRequestNotSupport("%d 权限信息不存在", userInfo.UserId)
	}

	req.Id = userInfo.UserId

	return s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: req.GetId(),
	})
}
