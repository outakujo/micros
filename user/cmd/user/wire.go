//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/data"
	"user/internal/server"
	"user/internal/service"
)

func wireApp(*conf.Server, *conf.Data, *conf.Auth, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet, newDiscovery, newApp))
}
