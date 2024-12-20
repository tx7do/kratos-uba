// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	"kratos-uba/app/core/service/internal/data"
	"kratos-uba/app/core/service/internal/server"
	"kratos-uba/app/core/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *v1.Bootstrap) (*kratos.App, func(), error) {
	entClient := data.NewEntClient(bootstrap, logger)
	client := data.NewRedisClient(bootstrap, logger)
	dataData, cleanup, err := data.NewData(entClient, client, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userService := service.NewUserService(logger, userRepo)
	applicationRepo := data.NewApplicationRepo(dataData, logger)
	applicationService := service.NewApplicationService(logger, applicationRepo)
	grpcServer := server.NewGRPCServer(bootstrap, logger, userService, applicationService)
	app := newApp(logger, registrar, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
