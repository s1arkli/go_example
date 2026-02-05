package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example/7.grpc/pb" // 根据你的 go.mod 调整
)

func main() {
	// 1. 建立连接
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 2. 创建客户端
	client := pb.NewUserSvcClient(conn)

	// 3. 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 4. 调用 RPC
	resp, err := client.GetUserInfo(ctx, &pb.UserReq{Id: 123})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	// 5. 打印结果
	log.Printf("响应: Name=%s, Age=%d, Gender=%d, Status=%v",
		resp.GetName(),
		resp.GetAge(),
		resp.GetGender(),
		resp.GetStatus(),
	)
}
