package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
	adminV1 "kratos-uba/gen/api/go/admin/service/v1"
	userV1 "kratos-uba/gen/api/go/user/service/v1"

	"kratos-uba/pkg/middleware/auth"
)

type UserService struct {
	adminV1.UserServiceHTTPServer

	log *log.Helper

	uc userV1.UserServiceClient
}

func NewUserService(uc userV1.UserServiceClient, logger log.Logger) *UserService {
	l := log.NewHelper(log.With(logger, "module", "user/service/admin-service"))
	return &UserService{
		log: l,
		uc:  uc,
	}
}

func (s *UserService) GetMe(ctx context.Context, _ *emptypb.Empty) (*userV1.User, error) {
	userInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, adminV1.ErrorRequestNotSupport("%d 权限信息不存在", userInfo.UserId)
	}

	return s.uc.GetUser(ctx, &userV1.GetUserRequest{
		Id: userInfo.UserId,
	})
}

func (s *UserService) ListUser(ctx context.Context, req *pagination.PagingRequest) (*userV1.ListUserResponse, error) {
	return s.uc.ListUser(ctx, req)
}

func (s *UserService) GetUser(ctx context.Context, req *userV1.GetUserRequest) (*userV1.User, error) {
	return s.uc.GetUser(ctx, req)
}

func (s *UserService) GetUserByUserName(ctx context.Context, req *userV1.GetUserByUserNameRequest) (*userV1.User, error) {
	return s.uc.GetUserByUserName(ctx, req)
}

func (s *UserService) CreateUser(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.User, error) {
	userInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, adminV1.ErrorRequestNotSupport("%d 权限信息不存在", userInfo.UserId)
	}

	if req.User == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	authority := "CUSTOMER_USER"

	req.OperatorId = userInfo.UserId
	req.User.CreatorId = &userInfo.UserId
	if req.User.Authority == nil {
		req.User.Authority = &authority
	}

	return s.uc.CreateUser(ctx, req)
}

func (s *UserService) UpdateUser(ctx context.Context, req *userV1.UpdateUserRequest) (*userV1.User, error) {
	userInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, adminV1.ErrorRequestNotSupport("%d 权限信息不存在", userInfo.UserId)
	}

	if req.User == nil {
		return nil, adminV1.ErrorBadRequest("错误的参数")
	}

	req.OperatorId = userInfo.UserId

	return s.uc.UpdateUser(ctx, req)
}

func (s *UserService) DeleteUser(ctx context.Context, req *userV1.DeleteUserRequest) (*emptypb.Empty, error) {
	userInfo, err := auth.FromContext(ctx)
	if err != nil {
		return nil, adminV1.ErrorRequestNotSupport("%d 权限信息不存在", userInfo.UserId)
	}

	req.OperatorId = userInfo.UserId

	return s.uc.DeleteUser(ctx, req)
}

func (s *UserService) UserExists(ctx context.Context, req *userV1.UserExistsRequest) (*userV1.UserExistsResponse, error) {
	return s.uc.UserExists(ctx, req)
}
