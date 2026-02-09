package main

import (
	"context"
	"fmt"
	"time"

	"example/7.grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewUserSvcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, _ := client.GetUserInfo(ctx, &pb.UserReq{Id: 123})

	fmt.Println(resp)
}
