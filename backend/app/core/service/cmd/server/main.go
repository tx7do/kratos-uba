package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"kratos-uba/pkg/service"
)

// go build -ldflags "-X main.Service.Version=x.y.z"

var version string

func newApp(ll log.Logger, rr registry.Registrar, gs *grpc.Server) *kratos.App {
	return bootstrap.NewApp(ll, rr, gs)
}

func main() {
	bootstrap.Bootstrap(initApp, trans.Ptr(service.CoreService), trans.Ptr(version))
}
