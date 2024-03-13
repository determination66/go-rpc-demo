package main

import (
	"go_rpc_demo/handler"
	"go_rpc_demo/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	// 实例化server
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic("net listen err")
	}
	// 注册handler逻辑
	err = server_proxy.RegisterHelloService(&handler.HelloService{})
	if err != nil {
		panic("server_proxy.RegisterHelloService failed!")
	}

	for {
		// 启动服务
		conn, err := listener.Accept()
		if err != nil {
			panic("accept err")
		}
		go rpc.ServeConn(conn)
	}

}
