package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-uba/app/admin/service/internal/service"

	adminV1 "kratos-uba/api/gen/go/admin/service/v1"

	"kratos-uba/pkg/middleware/auth"
)

// NewWhiteListMatcher 创建jwt白名单
func newRestWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]bool)
	whiteList[adminV1.OperationAuthenticationServiceLogin] = true
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func newRestMiddleware(authenticator authnEngine.Authenticator, authorizer authzEngine.Engine, logger log.Logger) []middleware.Middleware {
	var ms []middleware.Middleware
	ms = append(ms, logging.Server(logger))
	ms = append(ms, selector.Server(
		authn.Server(authenticator),
		auth.Server(),
		authz.Server(authorizer),
	).Match(newRestWhiteListMatcher()).Build())
	return ms
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	cfg *conf.Bootstrap, logger log.Logger,
	authenticator authnEngine.Authenticator, authorizer authzEngine.Engine,
	userSvc *service.UserService,
	authSvc *service.AuthenticationService,
	appSvc *service.ApplicationService,
) *http.Server {
	srv := rpc.CreateRestServer(cfg, newRestMiddleware(authenticator, authorizer, logger)...)

	adminV1.RegisterUserServiceHTTPServer(srv, userSvc)
	adminV1.RegisterAuthenticationServiceHTTPServer(srv, authSvc)
	adminV1.RegisterApplicationServiceHTTPServer(srv, appSvc)

	return srv
}
