package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/tx7do/kratos-bootstrap/rpc"

	"kratos-uba/app/core/service/internal/service"

	userV1 "kratos-uba/api/gen/go/user/service/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	cfg *conf.Bootstrap, logger log.Logger,
	userSvc *service.UserService,
	appSvc *service.ApplicationService,
) *grpc.Server {
	srv := rpc.CreateGrpcServer(cfg, logging.Server(logger))

	userV1.RegisterUserServiceServer(srv, userSvc)
	userV1.RegisterApplicationServiceServer(srv, appSvc)

	return srv
}
