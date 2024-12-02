package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-uba/pkg/service"
)

var version string

// go build -ldflags "-X main.Service.Version=x.y.z"

func newApp(ll log.Logger, rr registry.Registrar, ks *kafka.Server) *kratos.App {
	return bootstrap.NewApp(ll, rr, ks)
}

func main() {
	bootstrap.Bootstrap(initApp, trans.Ptr(service.LoggerService), trans.Ptr(version))
}
