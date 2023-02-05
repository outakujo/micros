package service

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"goods/internal/biz"

	pb "goods/api/v1"
)

type GoodsService struct {
	pb.UnimplementedGoodsServer
	biz *biz.GoodsBiz
	dis *etcd.Registry
}

func NewGoodsService(biz *biz.GoodsBiz, dis *etcd.Registry) *GoodsService {
	return &GoodsService{biz: biz, dis: dis}
}

func (s *GoodsService) CreateGoods(ctx context.Context, req *pb.CreateGoodsRequest) (*pb.Model, error) {
	goods := biz.Goods{
		Name:  req.Name,
		Price: req.Price,
		Count: req.Count,
	}
	create, err := s.biz.Create(ctx, goods)
	if err != nil {
		return nil, err
	}
	return &pb.Model{
		Id:    int64(create.ID),
		Name:  create.Name,
		Price: create.Price,
		Count: create.Count,
	}, nil
}
func (s *GoodsService) UpdateGoods(ctx context.Context, req *pb.UpdateGoodsRequest) (*pb.Model, error) {
	return &pb.Model{}, nil
}
func (s *GoodsService) DeleteGoods(ctx context.Context, req *pb.DeleteGoodsRequest) (*pb.DeleteGoodsReply, error) {
	err := s.biz.Delete(ctx, int(req.Id))
	return &pb.DeleteGoodsReply{}, err
}
func (s *GoodsService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.Model, error) {
	get, err := s.biz.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Model{
		Id:    int64(get.ID),
		Name:  get.Name,
		Price: get.Price,
		Count: get.Count,
	}, nil
}
func (s *GoodsService) ListGoods(ctx context.Context, req *pb.ListGoodsRequest) (*pb.ListGoodsReply, error) {
	p, err := s.biz.List(ctx, int(req.Page), int(req.Size))
	if err != nil {
		return nil, err
	}
	models := make([]*pb.Model, 0)
	for _, g := range p.List {
		m := &pb.Model{
			Id:    int64(g.ID),
			Name:  g.Name,
			Price: g.Price,
			Count: g.Count,
		}
		models = append(models, m)
	}
	return &pb.ListGoodsReply{
		List:    models,
		MaxPage: p.MaxPage,
		Count:   p.Count,
	}, nil
}
