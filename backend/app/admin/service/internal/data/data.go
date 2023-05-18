package data

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	userV1 "kratos-bi/gen/api/go/user/service/v1"

	"kratos-bi/gen/api/go/common/conf"
	"kratos-bi/pkg/bootstrap"
	"kratos-bi/pkg/service"
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
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Data.Redis.Addr,
		Password:     cfg.Data.Redis.Password,
		DB:           int(cfg.Data.Redis.Db),
		DialTimeout:  cfg.Data.Redis.DialTimeout.AsDuration(),
		WriteTimeout: cfg.Data.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  cfg.Data.Redis.ReadTimeout.AsDuration(),
	})
	if rdb == nil {
		l.Fatalf("failed opening connection to redis")
	}
	rdb.AddHook(redisotel.NewTracingHook())

	return rdb
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bootstrap.NewConsulRegistry(cfg.Registry)
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

func NewUserServiceClient(r registry.Discovery, cfg *conf.Bootstrap) userV1.UserServiceClient {
	return userV1.NewUserServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, cfg))
}

func NewApplicationServiceClient(r registry.Discovery, cfg *conf.Bootstrap) userV1.ApplicationServiceClient {
	return userV1.NewApplicationServiceClient(bootstrap.CreateGrpcClient(context.Background(), r, service.CoreService, cfg))
}
