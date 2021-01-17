package grpcExample

import (
	"context"
	"fmt"
	pb "golangStudy/grpcExample/go/protos"
	"google.golang.org/grpc"

)

func test2() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接异常", err)
	}

	defer conn.Close()
	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)

	req.Name = "sssss"
	//基于如上需求，context包应用而生。context包可以提供一个请求从API请求边界到各goroutine的请求域数据传递、取消信号及截至时间等能力。详细原理请看下文。
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常", err)
	}
	fmt.Println("结果:", resp)
}