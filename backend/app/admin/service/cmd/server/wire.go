//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"kratos-bi/app/admin/service/internal/biz"
	"kratos-bi/app/admin/service/internal/data"
	"kratos-bi/app/admin/service/internal/server"
	"kratos-bi/app/admin/service/internal/service"

	"kratos-bi/gen/api/go/common/conf"
)

// initApp init kratos application.
func initApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
