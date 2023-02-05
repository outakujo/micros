package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	err := data.DB.AutoMigrate(biz.User{})
	if err != nil {
		panic(err)
	}
	return &UserRepo{data: data, log: log.NewHelper(logger)}
}

func (u *UserRepo) Create(ctx context.Context, user biz.User) (biz.User, error) {
	err := u.data.Create(&user).Error
	return user, err
}

func (u *UserRepo) Update(ctx context.Context, user biz.User) (biz.User, error) {
	err := u.data.Save(&user).Error
	return user, err
}

func (u *UserRepo) Delete(ctx context.Context, username string) error {
	err := u.data.Where("username=?", username).Delete(biz.User{}).Error
	return err
}

func (u *UserRepo) Get(ctx context.Context, username string) (biz.User, error) {
	var user biz.User
	err := u.data.Model(biz.User{}).Where("username=?", username).First(&user).Error
	return user, err
}

func (u *UserRepo) List(ctx context.Context, page, size int) (biz.PageUser, error) {
	pu := biz.PageUser{List: make([]*biz.User, 0)}
	db := u.data.DB.Model(biz.User{}).Order("id desc")
	helper := NewPageHelper(db, int64(page), int64(size))
	err := helper.List(&pu.List)
	if err != nil {
		return pu, err
	}
	pu.Count = helper.Count
	pu.MaxPage = helper.MaxPage
	return pu, nil
}
