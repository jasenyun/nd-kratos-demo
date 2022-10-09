package service

import (
	"context"
	v1 "jasenServer/api/helloworld/v1"
	"jasenServer/internal/biz"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) CreateGreeter(ctx context.Context, in *v1.CreateRequest) (*v1.CreateReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Name: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.CreateReply{Message: "Hello " + g.Name}, nil
}
func (s *GreeterService) FindById(ctx context.Context, in *v1.FindRequest) (*v1.FindReply, error) {
	g, err := s.uc.FindById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.FindReply{Id: g.Id, Name: g.Name}, nil
}
