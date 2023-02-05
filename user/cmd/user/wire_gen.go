// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/data"
	"user/internal/server"
	"user/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(c *conf.Server, da *conf.Data, log2 log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(da, log2)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, log2)
	userBiz := biz.NewUserBiz(userRepo, log2)
	registry := newDiscovery(da)
	userService := service.NewUserService(userBiz, registry, log2)
	grpcServer := server.NewGRPCServer(c, userService, log2)
	httpServer := server.NewHTTPServer(c, userService, log2)
	app := newApp(log2, grpcServer, httpServer, registry)
	return app, func() {
		cleanup()
	}, nil
}
