package main

import (
	"context"
	"go_rpc_demo/protoc/hello_world"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	hello_world.GreeterServer
}

//SayHello(context.Context, *HelloRequest) (*HelloReply, error)

// SayHello proto可以生成go语言
func (s *Server) SayHello(ctx context.Context, request *hello_world.HelloRequest) (
	reply *hello_world.HelloReply, err error) {
	return &hello_world.HelloReply{
		Message: "hello," + request.Name,
	}, nil
}

func main() {
	grpcServer := grpc.NewServer()

	hello_world.RegisterGreeterServer(grpcServer, &Server{})

	// 监听 gRPC 服务器
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 开始接受客户端连接并处理 gRPC 请求
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
