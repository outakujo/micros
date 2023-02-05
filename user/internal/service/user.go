package service

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"user/internal/biz"

	pb "user/api/v1"
	v1 "user/api/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	biz *biz.UserBiz
	dis *etcd.Registry
	log *log.Helper
}

func NewUserService(biz *biz.UserBiz, dis *etcd.Registry, logger log.Logger) *UserService {
	return &UserService{biz: biz, dis: dis, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.Model, error) {
	var u biz.User
	u.Username = req.Username
	u.Nickname = req.Nickname
	u.Password = req.Password
	create, err := s.biz.Create(ctx, u)
	if err != nil {
		return nil, err
	}
	return &pb.Model{
		Username: create.Username,
		Nickname: create.Nickname,
		Id:       int64(create.ID),
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Model, error) {
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	client, err := goodsClient(s.dis)
	if err != nil {
		return nil, err
	}
	getGoods, err := client.GetGoods(timeout, &v1.GetGoodsRequest{Id: 1})
	if err != nil {
		return nil, err
	}
	return &pb.Model{
		Username: getGoods.Name,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.biz.Delete(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.Model, error) {
	get, err := s.biz.Get(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	return &pb.Model{
		Username: get.Username,
		Nickname: get.Nickname,
		Id:       int64(get.ID),
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	p, err := s.biz.List(ctx, int(req.Page), int(req.Size))
	if err != nil {
		return nil, err
	}
	models := make([]*pb.Model, 0)
	for _, user := range p.List {
		m := &pb.Model{
			Id:       int64(user.ID),
			Username: user.Username,
			Nickname: user.Nickname,
		}
		models = append(models, m)
	}
	return &pb.ListUserReply{
		List:    models,
		MaxPage: p.MaxPage,
		Count:   p.Count,
	}, nil
}
