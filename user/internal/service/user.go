package service

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"time"
	"user/internal/biz"
	"user/internal/conf"

	v1 "user/api/v1"
)

type UserService struct {
	v1.UnimplementedUserServer
	biz *biz.UserBiz
	dis *etcd.Registry
	log *log.Helper
	au  *conf.Auth
}

func NewUserService(biz *biz.UserBiz, dis *etcd.Registry, au *conf.Auth, logger log.Logger) *UserService {
	return &UserService{biz: biz, dis: dis, log: log.NewHelper(logger), au: au}
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.Model, error) {
	if md, ok := metadata.FromServerContext(ctx); ok {
		// 从header里token值
		extra := md.Get("token")
		s.log.Infof("extra %s", extra)
	}
	var u biz.User
	u.Username = req.Username
	u.Nickname = req.Nickname
	u.Password = req.Password
	create, err := s.biz.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return &v1.Model{
		Username: create.Username,
		Nickname: create.Nickname,
		Id:       int64(create.ID),
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.Model, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	client, err := goodsClient(s.dis, s.au)
	if err != nil {
		return nil, err
	}
	getGoods, err := client.GetGoods(timeout, &v1.GetGoodsRequest{Id: 1})
	if err != nil {
		return nil, err
	}
	return &v1.Model{
		Username: getGoods.Name,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := s.biz.Delete(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.Model, error) {
	get, err := s.biz.Get(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &v1.Model{
		Username: get.Username,
		Nickname: get.Nickname,
		Id:       int64(get.ID),
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	p, err := s.biz.List(ctx, int(req.Page), int(req.Size))
	if err != nil {
		return nil, err
	}
	models := make([]*v1.Model, 0)
	for _, user := range p.List {
		m := &v1.Model{
			Id:       int64(user.ID),
			Username: user.Username,
			Nickname: user.Nickname,
		}
		models = append(models, m)
	}
	return &v1.ListUserReply{
		List:    models,
		MaxPage: p.MaxPage,
		Count:   p.Count,
	}, nil
}
