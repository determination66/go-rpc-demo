package main

import (
	"context"
	"fmt"
	"go_rpc_demo/protoc/hello_world"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	hello_world.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *hello_world.HelloRequest) (*hello_world.HelloReply, error) {
	fmt.Println("成功SayHello")

	time.Sleep(4 * time.Second)

	reply := &hello_world.HelloReply{
		Message: "hello," + req.Name,
	}

	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()

	hello_world.RegisterGreeterServer(grpcServer, &server{})

	// 监听 gRPC 服务器
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 开始接受客户端连接并处理 gRPC 请求
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
