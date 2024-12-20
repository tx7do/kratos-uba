package data

import (
	"context"

	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/kafka"

	"github.com/redis/go-redis/v9"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	userV1 "kratos-uba/api/gen/go/user/service/v1"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"
	bRegistry "github.com/tx7do/kratos-bootstrap/registry"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-uba/pkg/service"
)

// Data .
type Data struct {
	log *log.Helper
	rdb *redis.Client

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine

	appClient userV1.ApplicationServiceClient
}

// NewData .
func NewData(
	logger log.Logger,
	redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	appClient userV1.ApplicationServiceClient,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/agent-service"))

	d := &Data{
		log:           l,
		authenticator: authenticator,
		authorizer:    authorizer,
		rdb:           redisClient,
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

func NewApplicationServiceClient(r registry.Discovery, cfg *conf.Bootstrap) userV1.ApplicationServiceClient {
	return userV1.NewApplicationServiceClient(rpc.CreateGrpcClient(context.Background(), r, service.CoreService, cfg))
}

func NewKafkaBroker(cfg *conf.Bootstrap) broker.Broker {
	b := kafka.NewBroker(
		broker.WithAddress(cfg.Data.Kafka.Endpoints...),
		broker.WithCodec(cfg.Data.Kafka.Codec),
	)
	if b == nil {
		return nil
	}

	_ = b.Init()

	if err := b.Connect(); err != nil {
		return nil
	}

	return b
}
