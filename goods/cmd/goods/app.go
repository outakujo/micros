package main

import (
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"go.etcd.io/etcd/client/v3"
	"goods/internal/conf"
)

func newDiscovery(cd *conf.Data) *etcd.Registry {
	c := cd.Etcd
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{fmt.Sprintf("%v:%v", c.Ip, c.Port)},
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(client)
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, dis *etcd.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(dis),
	)
}
