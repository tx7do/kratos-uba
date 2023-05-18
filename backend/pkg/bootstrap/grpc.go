package bootstrap

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	midRateLimit "github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	kratosGrpc "github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-bi/gen/api/go/common/conf"
)

const defaultTimeout = 5 * time.Second

// CreateGrpcClient 创建GRPC客户端
func CreateGrpcClient(ctx context.Context, r registry.Discovery, serviceName string, cfg *conf.Bootstrap, m ...middleware.Middleware) grpc.ClientConnInterface {
	endpoint := "discovery:///" + serviceName

	var ms []middleware.Middleware
	timeout := defaultTimeout
	if cfg.Client != nil && cfg.Client.Grpc != nil {
		if cfg.Client.Grpc.Timeout != nil {
			timeout = cfg.Client.Grpc.Timeout.AsDuration()
		}

		if cfg.Client.Grpc.Middleware != nil {
			if cfg.Client.Grpc.Middleware.GetEnableRecovery() {
				ms = append(ms, recovery.Recovery())
			}
			if cfg.Client.Grpc.Middleware.GetEnableTracing() {
				ms = append(ms, tracing.Client())
			}
			if cfg.Client.Grpc.Middleware.GetEnableValidate() {
				ms = append(ms, validate.Validator())
			}
		}
	}
	ms = append(ms, m...)

	conn, err := kratosGrpc.DialInsecure(
		ctx,
		kratosGrpc.WithEndpoint(endpoint),
		kratosGrpc.WithDiscovery(r),
		kratosGrpc.WithTimeout(timeout),
		kratosGrpc.WithMiddleware(ms...),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}

	return conn
}

// CreateGrpcServer 创建GRPC服务端
func CreateGrpcServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosGrpc.Server {
	var opts []kratosGrpc.ServerOption

	var ms []middleware.Middleware
	if cfg.Server != nil && cfg.Server.Grpc != nil && cfg.Server.Grpc.Middleware != nil {
		if cfg.Server.Grpc.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Server.Grpc.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Server())
		}
		if cfg.Server.Grpc.Middleware.GetEnableValidate() {
			ms = append(ms, validate.Validator())
		}
		if cfg.Server.Grpc.Middleware.GetEnableCircuitBreaker() {
		}
		if cfg.Server.Grpc.Middleware.Limiter != nil {
			var limiter ratelimit.Limiter
			switch cfg.Server.Grpc.Middleware.Limiter.GetName() {
			case "bbr":
				limiter = bbr.NewLimiter()
			}
			ms = append(ms, midRateLimit.Server(midRateLimit.WithLimiter(limiter)))
		}
	}
	ms = append(ms, m...)
	opts = append(opts, kratosGrpc.Middleware(ms...))

	if cfg.Server.Grpc.Network != "" {
		opts = append(opts, kratosGrpc.Network(cfg.Server.Grpc.Network))
	}
	if cfg.Server.Grpc.Addr != "" {
		opts = append(opts, kratosGrpc.Address(cfg.Server.Grpc.Addr))
	}
	if cfg.Server.Grpc.Timeout != nil {
		opts = append(opts, kratosGrpc.Timeout(cfg.Server.Grpc.Timeout.AsDuration()))
	}

	return kratosGrpc.NewServer(opts...)
}
