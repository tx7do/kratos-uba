package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-uba/app/report/service/internal/service"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	v1 "kratos-uba/api/gen/go/report/service/v1"

	"github.com/tx7do/kratos-bootstrap/rpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	cfg *conf.Bootstrap, logger log.Logger,
	reportService *service.ReportService,
) *grpc.Server {
	srv := rpc.CreateGrpcServer(cfg, logging.Server(logger))

	v1.RegisterReportServiceServer(srv, reportService)

	return srv
}
