package main

import (
	"fmt"
	"go_rpc_demo/global"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(p *global.Person, res *global.Res) error {
	res.Code = 200
	res.Msg = "hello," + p.Name

	fmt.Println("server:", p)
	return nil
}

// 服务端rpc代码
func main() {
	// 实例化server
	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic("net listen err")
	}
	// 注册handler逻辑
	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		panic("rpc register failed!")
	}

	for {
		// 启动服务
		conn, err := listener.Accept()
		if err != nil {
			panic("accept err")
		}
		rpc.ServeConn(conn)
	}

}
