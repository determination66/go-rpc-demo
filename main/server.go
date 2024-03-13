package main

import (
	"fmt"
	"go_rpc_demo/global"
	"go_rpc_demo/handler"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(p *global.Person, res *global.Res) error {
	res.Code = 200
	res.Msg = "hello," + p.Name
	fmt.Println(p)
	return nil
}

func main() {
	// 实例化server
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic("net listen err")
	}
	// 注册handler逻辑
	err = rpc.RegisterName(handler.HelloServiceName, &HelloService{})
	if err != nil {
		panic("rpc register failed!")
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
