package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"example/7.grpc/pb"
)

type userServer struct {
	pb.UnimplementedUserSvcServer
}

func main() {
	listener, _ := net.Listen("tcp", ":50051")

	server := grpc.NewServer()

	pb.RegisterUserSvcServer(server, &userServer{})

	fmt.Println("grpc server listening on :50051")
	server.Serve(listener)
}

func (u *userServer) GetUserInfo(ctx context.Context, req *pb.UserReq) (*pb.User, error) {
	return &pb.User{
		Age:    25,
		Gender: 1,
		Name:   "张三",
		Status: pb.Status_STATUS_UNSPECIFIED,
	}, nil
}
