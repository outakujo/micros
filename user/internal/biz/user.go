package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserBiz struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserBiz(repo UserRepo, logger log.Logger) *UserBiz {
	return &UserBiz{repo: repo, log: log.NewHelper(logger)}
}

type UserRepo interface {
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) (User, error)
	List(ctx context.Context, page, size int) (PageUser, error)
}

func (r *UserBiz) Create(ctx context.Context, user User) (User, error) {
	r.log.WithContext(ctx).Infof("create %v", user)
	return r.repo.Create(ctx, user)
}

func (r *UserBiz) Get(ctx context.Context, username string) (User, error) {
	r.log.WithContext(ctx).Infof("get %v", username)
	return r.repo.Get(ctx, username)
}

func (r *UserBiz) Delete(ctx context.Context, username string) error {
	r.log.WithContext(ctx).Infof("delete %v", username)
	return r.repo.Delete(ctx, username)
}

func (r *UserBiz) List(ctx context.Context, page, size int) (PageUser, error) {
	return r.repo.List(ctx, page, size)
}
