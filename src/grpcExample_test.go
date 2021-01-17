package main


import (
	pb "golangStudy/grpcExample/go/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"testing"
	grpcExample "golangStudy/grpcExample/go"
)


func TestAdd3(t *testing.T) {
	fmt.Println(234)
	if ans := grpcExample.Add3(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}

	//grpcExample.Add3(1, 2)
	//var u = grpcExample.UserInfoService{}
		//resp, err := client.GetUserInfo(context.Background(), req)

	 //resp, err := u.GetUserInfo( req *pb.UserRequest)

}