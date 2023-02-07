package service

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"time"
	v1 "user/api/v1"
	"user/internal/conf"
)

func goodsClient(dis *etcd.Registry, a *conf.Auth) (v1.GoodsClient, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	endpoint := "discovery:///goods"
	selector.SetGlobalSelector(wrr.NewBuilder())
	dial, err := grpc.DialInsecure(timeout, grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis),
		grpc.WithNodeFilter(filter.Version("1.0")),
		grpc.WithMiddleware(circuitbreaker.Client(),
			jwt.Client(func(token *jwt2.Token) (interface{}, error) {
				return []byte(a.GoodsKey), nil
			})),
	)
	if err != nil {
		return nil, err
	}
	return v1.NewGoodsClient(dial), nil
}
