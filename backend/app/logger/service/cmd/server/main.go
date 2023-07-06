package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-bi/pkg/bootstrap"
	"kratos-bi/pkg/service"
)

// go build -ldflags "-X main.Service.Version=x.y.z"

var (
	Service = bootstrap.NewServiceInfo(
		service.LoggerService,
		"1.0.0",
		"",
	)
)

func newApp(ll log.Logger, rr registry.Registrar, ks *kafka.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(ll),
		kratos.Server(
			ks,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	cfg, ll, reg := bootstrap.Bootstrap(Service)

	app, cleanup, err := initApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
