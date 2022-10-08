package service

import (
	"context"

	v1 "jasenServer/api/helloworld/v1"
	"jasenServer/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Name: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Name}, nil
}

func (s *GreeterService) FindById(ctx context.Context, in *v1.IdRequest) (*v1.IdReply, error) {
	g, err := s.uc.FindById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.IdReply{Id: g.Id, Name: g.Name}, nil
}
