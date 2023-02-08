package service

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"goods/internal/biz"

	v1 "goods/api/v1"
)

type GoodsService struct {
	v1.UnimplementedGoodsServer
	biz *biz.GoodsBiz
	dis *etcd.Registry
}

func NewGoodsService(biz *biz.GoodsBiz, dis *etcd.Registry) *GoodsService {
	return &GoodsService{biz: biz, dis: dis}
}

func (s *GoodsService) CreateGoods(ctx context.Context, req *v1.CreateGoodsRequest) (*v1.Model, error) {
	goods := biz.Goods{
		Name:  req.Name,
		Price: req.Price,
		Count: req.Count,
	}
	create, err := s.biz.Create(ctx, goods)
	if err != nil {
		return nil, err
	}
	return &v1.Model{
		Id:    int64(create.ID),
		Name:  create.Name,
		Price: create.Price,
		Count: create.Count,
	}, nil
}
func (s *GoodsService) UpdateGoods(ctx context.Context, req *v1.UpdateGoodsRequest) (*v1.Model, error) {
	return &v1.Model{}, nil
}
func (s *GoodsService) DeleteGoods(ctx context.Context, req *v1.DeleteGoodsRequest) (*v1.DeleteGoodsReply, error) {
	err := s.biz.Delete(ctx, int(req.Id))
	return &v1.DeleteGoodsReply{}, err
}
func (s *GoodsService) GetGoods(ctx context.Context, req *v1.GetGoodsRequest) (*v1.Model, error) {
	get, err := s.biz.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &v1.Model{
		Id:    int64(get.ID),
		Name:  get.Name,
		Price: get.Price,
		Count: get.Count,
	}, nil
}
func (s *GoodsService) ListGoods(ctx context.Context, req *v1.ListGoodsRequest) (*v1.ListGoodsReply, error) {
	p, err := s.biz.List(ctx, int(req.Page), int(req.Size))
	if err != nil {
		return nil, err
	}
	models := make([]*v1.Model, 0)
	for _, g := range p.List {
		m := &v1.Model{
			Id:    int64(g.ID),
			Name:  g.Name,
			Price: g.Price,
			Count: g.Count,
		}
		models = append(models, m)
	}
	return &v1.ListGoodsReply{
		List:    models,
		MaxPage: p.MaxPage,
		Count:   p.Count,
	}, nil
}
