package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
)

type GoodsRepo struct {
	data *Data
	log  *log.Helper
}

func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	err := data.DB.AutoMigrate(biz.Goods{})
	if err != nil {
		panic(err)
	}
	return &GoodsRepo{data: data, log: log.NewHelper(logger)}
}

func (u *GoodsRepo) Create(ctx context.Context, Goods biz.Goods) (biz.Goods, error) {
	err := u.data.Create(&Goods).Error
	return Goods, err
}

func (u *GoodsRepo) Update(ctx context.Context, Goods biz.Goods) (biz.Goods, error) {
	err := u.data.Save(&Goods).Error
	return Goods, err
}

func (u *GoodsRepo) Delete(ctx context.Context, id int) error {
	err := u.data.Where("id=?", id).Delete(biz.Goods{}).Error
	return err
}

func (u *GoodsRepo) Get(ctx context.Context, id int) (biz.Goods, error) {
	var Goods biz.Goods
	err := u.data.Model(biz.Goods{}).Where("id=?", id).First(&Goods).Error
	return Goods, err
}

func (u *GoodsRepo) List(ctx context.Context, page, size int) (biz.PageGoods, error) {
	pu := biz.PageGoods{List: make([]*biz.Goods, 0)}
	db := u.data.DB.Model(biz.Goods{}).Order("id desc")
	helper := NewPageHelper(db, int64(page), int64(size))
	err := helper.List(&pu.List)
	if err != nil {
		return pu, err
	}
	pu.Count = helper.Count
	pu.MaxPage = helper.MaxPage
	return pu, nil
}
