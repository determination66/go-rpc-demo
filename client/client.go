package main

import (
	"context"
	"fmt"
	"go_rpc_demo/protoc/hello_world"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 连接 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	// 创建 Greeter 客户端
	client := hello_world.NewGreeterClient(conn)

	// 调用 SayHello 方法
	resp, err := client.SayHello(context.Background(), &hello_world.HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	// 打印响应
	fmt.Println("Response:", resp.Message)
}
