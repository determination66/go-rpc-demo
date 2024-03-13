package main

import (
	"fmt"
	"go_rpc_demo/global"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		panic("rpc.RegisterName:" + err.Error())
	}

	listener, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic("net.Listen:" + err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic("listener.Accept():" + err.Error())
		}
		go jsonrpc.ServeConn(conn)
	}

}
