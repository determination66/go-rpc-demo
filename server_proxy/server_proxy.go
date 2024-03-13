package server_proxy

import (
	"go_rpc_demo/global"
	"go_rpc_demo/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(*global.Person, *global.Res) error
}

// RegisterHelloService 注册Service
func RegisterHelloService(svc HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, svc)
}
