package service

import (
	"context"

	"kratos-layout/internal/biz"
	pb "kratos-layout/proto/greeter"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) CreateGreeter(ctx context.Context, req *pb.CreateGreeterRequest) (*pb.CreateGreeterReply, error) {
	return &pb.CreateGreeterReply{}, nil
}
func (s *GreeterService) UpdateGreeter(ctx context.Context, req *pb.UpdateGreeterRequest) (*pb.UpdateGreeterReply, error) {
	return &pb.UpdateGreeterReply{}, nil
}
func (s *GreeterService) DeleteGreeter(ctx context.Context, req *pb.DeleteGreeterRequest) (*pb.DeleteGreeterReply, error) {
	return &pb.DeleteGreeterReply{}, nil
}
func (s *GreeterService) GetGreeter(ctx context.Context, req *pb.GetGreeterRequest) (*pb.GetGreeterReply, error) {
	return &pb.GetGreeterReply{}, nil
}
func (s *GreeterService) ListGreeter(ctx context.Context, req *pb.ListGreeterRequest) (*pb.ListGreeterReply, error) {
	return &pb.ListGreeterReply{}, nil
}
func (s *GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}

	return &pb.HelloReply{Message: "Hello " + g.Hello}, nil
}
