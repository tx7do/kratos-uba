package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	authV1 "go-wind-uba/api/gen/go/authentication/service/v1"

	"go-wind-uba/pkg/constants"

	"go-wind-uba/pkg/middleware/auth"
)

// requirePlatformAdmin 把通知域的平台级配置与台账挡在平台管理员之内。
//
// 渠道/规则/模板/台账四张表无租户列（平台全局资源），租户侧可读即全平台泄漏。
// 鉴权发生在 admin BFF 边缘（JWT 中间件已将载荷写入 context），
// core 侧 gRPC 是内网信任模型、不再重复校验——与 uba 其余管理域一致。
func requirePlatformAdmin(ctx context.Context, log *log.Helper, resource string) error {
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return authV1.ErrorUnauthorized("authentication required: %v", err)
	}

	// 平台管理员判定走角色码（platform:admin）；JWT 载荷不含独立标志位
	for _, role := range operator.GetRoles() {
		if role == constants.PlatformAdminRoleCode {
			return nil
		}
	}

	log.Warnf("operator [%d] (tenant %d) is not a platform administrator, %s access denied",
		operator.GetUserId(), operator.GetTenantId(), resource)

	return authV1.ErrorForbidden("%s is platform-level configuration, only platform administrators may access it", resource)
}
