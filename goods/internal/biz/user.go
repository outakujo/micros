package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type GoodsBiz struct {
	repo GoodsRepo
	log  *log.Helper
}

func NewGoodsBiz(repo GoodsRepo, logger log.Logger) *GoodsBiz {
	return &GoodsBiz{repo: repo, log: log.NewHelper(logger)}
}

type GoodsRepo interface {
	Create(ctx context.Context, Goods Goods) (Goods, error)
	Update(ctx context.Context, Goods Goods) (Goods, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (Goods, error)
	List(ctx context.Context, page, size int) (PageGoods, error)
}

func (r *GoodsBiz) Create(ctx context.Context, Goods Goods) (Goods, error) {
	r.log.WithContext(ctx).Infof("create %v", Goods)
	return r.repo.Create(ctx, Goods)
}

func (r *GoodsBiz) Get(ctx context.Context, id int) (Goods, error) {
	r.log.WithContext(ctx).Infof("get %v", id)
	return r.repo.Get(ctx, id)
}

func (r *GoodsBiz) Delete(ctx context.Context, id int) error {
	r.log.WithContext(ctx).Infof("delete %v", id)
	return r.repo.Delete(ctx, id)
}

func (r *GoodsBiz) List(ctx context.Context, page, size int) (PageGoods, error) {
	return r.repo.List(ctx, page, size)
}
