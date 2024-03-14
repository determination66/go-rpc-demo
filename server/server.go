package main

import (
	"fmt"
	"go_rpc_demo/protoc/stream"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type server struct {
	stream.GreeterServer
}

// 服务端 单向流
func (s *server) GetStream(req *stream.StreamReqData, res stream.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		res.Send(&stream.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端 单向流
func (s *server) PutStream(cliStr stream.Greeter_PutStreamServer) error {

	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

// 客户端服务端 双向流
func (s *server) AllStream(allStr stream.Greeter_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data.Data)
		}
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&stream.StreamResData{Data: "ssss"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	grpcServer := grpc.NewServer()

	stream.RegisterGreeterServer(grpcServer, &server{})

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
