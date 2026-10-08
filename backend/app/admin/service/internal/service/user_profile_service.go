package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"google.golang.org/protobuf/types/known/emptypb"

	adminV1 "go-wind-uba/api/gen/go/admin/service/v1"
	authenticationV1 "go-wind-uba/api/gen/go/authentication/service/v1"
	identityV1 "go-wind-uba/api/gen/go/identity/service/v1"
	permissionV1 "go-wind-uba/api/gen/go/permission/service/v1"

	"go-wind-uba/pkg/middleware/auth"
)

type UserProfileService struct {
	adminV1.UserProfileServiceHTTPServer

	log *log.Helper

	userServiceClient     identityV1.UserServiceClient
	tenantServiceClient   identityV1.TenantServiceClient
	orgUnitServiceClient  identityV1.OrgUnitServiceClient
	positionServiceClient identityV1.PositionServiceClient

	roleServiceClient permissionV1.RoleServiceClient

	userCredentialServiceClient authenticationV1.UserCredentialServiceClient

	authenticationServiceClient authenticationV1.AuthenticationServiceClient
}

func NewUserProfileService(
	ctx *bootstrap.Context,
	userServiceClient identityV1.UserServiceClient,
	tenantServiceClient identityV1.TenantServiceClient,
	orgUnitServiceClient identityV1.OrgUnitServiceClient,
	positionServiceClient identityV1.PositionServiceClient,
	roleServiceClient permissionV1.RoleServiceClient,
	userCredentialServiceClient authenticationV1.UserCredentialServiceClient,
	authenticationServiceClient authenticationV1.AuthenticationServiceClient,
) *UserProfileService {
	return &UserProfileService{
		log:                         ctx.NewLoggerHelper("user-profile/service/admin-service"),
		userServiceClient:           userServiceClient,
		tenantServiceClient:         tenantServiceClient,
		orgUnitServiceClient:        orgUnitServiceClient,
		positionServiceClient:       positionServiceClient,
		roleServiceClient:           roleServiceClient,
		userCredentialServiceClient: userCredentialServiceClient,
		authenticationServiceClient: authenticationServiceClient,
	}
}

func (s *UserProfileService) GetUser(ctx context.Context, _ *emptypb.Empty) (*identityV1.User, error) {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := s.userServiceClient.Get(ctx, &identityV1.GetUserRequest{
		QueryBy: &identityV1.GetUserRequest_Id{
			Id: operator.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserProfileService) UpdateUser(ctx context.Context, req *identityV1.UpdateUserRequest) (*emptypb.Empty, error) {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.Id = trans.Ptr(operator.UserId)
	req.Id = operator.UserId

	return s.userServiceClient.Update(ctx, req)
}

func (s *UserProfileService) ChangePassword(ctx context.Context, req *identityV1.ChangePasswordRequest) (*emptypb.Empty, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	if _, err = s.userCredentialServiceClient.ChangeCredential(ctx, &authenticationV1.ChangeCredentialRequest{
		IdentityType:  authenticationV1.UserCredential_USERNAME,
		Identifier:    operator.GetUsername(),
		OldCredential: req.GetOldPassword(),
		NewCredential: req.GetNewPassword(),
	}); err != nil {
		return nil, err
	}

	// 改密成功后吊销本人全部客户端类型的令牌（含当前会话），
	// 防止凭据泄露后旧令牌继续可用；前端在成功回调中引导重新登录。
	for _, clientType := range []authenticationV1.ClientType{
		authenticationV1.ClientType_admin,
		authenticationV1.ClientType_collector,
	} {
		if _, rErr := s.authenticationServiceClient.Logout(ctx, &authenticationV1.LogoutRequest{
			UserId:     operator.GetUserId(),
			ClientType: clientType,
		}); rErr != nil {
			s.log.Errorf("revoke tokens after password change failed for user [%d] [%v]: %v",
				operator.GetUserId(), clientType, rErr)
		}
	}

	return &emptypb.Empty{}, nil
}
