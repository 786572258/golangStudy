package grpcExample

import (
	pb "golangStudy/grpcExample/go/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)


func Add3(num1 int, num2 int) int {
	if u == struct{}{} {

	}
	return num1 + num2
}

type UserInfoService struct {

}

var u = UserInfoService{}

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	resp = &pb.UserResponse {
		Id:1,
		Name:name,
	}
	err = nil
	return
}

func Main2() int{
	addr := "127.0.0.1:8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("监听异常", err)
	}
	fmt.Println("开始监听")
	s := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(s, &u)
	s.Serve(lis)
	return 1

}

