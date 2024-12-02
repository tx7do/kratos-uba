package data

import (
	"context"

	"github.com/redis/go-redis/v9"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"
	bRegistry "github.com/tx7do/kratos-bootstrap/registry"
	"github.com/tx7do/kratos-bootstrap/rpc"

	userV1 "kratos-uba/api/gen/go/user/service/v1"

	"kratos-uba/pkg/service"
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine

	userClient userV1.UserServiceClient
	appClient  userV1.ApplicationServiceClient
}

// NewData .
func NewData(
	logger log.Logger,
	redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	userClient userV1.UserServiceClient,
	appClient userV1.ApplicationServiceClient,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		rdb:           redisClient,
		log:           l,
		authenticator: authenticator,
		authorizer:    authorizer,
		userClient:    userClient,
		appClient:     appClient,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))
	return redisClient.NewClient(cfg.Data)
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bRegistry.NewDiscovery(cfg.Registry)
}

// NewAuthenticator 创建认证器
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Rest.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Rest.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer 创建权鉴器
func NewAuthorizer() authzEngine.Engine {
	return noop.State{}
}

func NewUserServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}

func NewApplicationServiceClient(r registry.Discovery, c *conf.Bootstrap) userV1.ApplicationServiceClient {
	return userV1.NewApplicationServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, c))
}
