package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/tx7do/go-utils/trans"
	"github.com/tx7do/kratos-bootstrap/bootstrap"

	"kratos-uba/pkg/service"
)

// go build -ldflags "-X main.Service.Version=x.y.z"

var version string

func newApp(ll log.Logger, rr registry.Registrar, hs *http.Server) *kratos.App {
	return bootstrap.NewApp(ll, rr, hs)
}

func main() {
	bootstrap.Bootstrap(initApp, trans.Ptr(service.AgentService), trans.Ptr(version))
}
