package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-uba/app/core/service/internal/service"

	"kratos-uba/gen/api/go/common/conf"
	userV1 "kratos-uba/gen/api/go/user/service/v1"

	"kratos-uba/pkg/bootstrap"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	cfg *conf.Bootstrap, logger log.Logger,
	userSvc *service.UserService,
	appSvc *service.ApplicationService,
) *grpc.Server {
	srv := bootstrap.CreateGrpcServer(cfg, logging.Server(logger))

	userV1.RegisterUserServiceServer(srv, userSvc)
	userV1.RegisterApplicationServiceServer(srv, appSvc)

	return srv
}
