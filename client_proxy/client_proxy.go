package client_proxy

import (
	"go_rpc_demo/global"
	"go_rpc_demo/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

// NewHelloServiceStub 创建Stub
func NewHelloServiceStub(network, address string) *HelloServiceStub {
	conn, err := rpc.Dial(network, address)
	if err != nil {
		panic("rpc.Dial:" + err.Error())
	}
	return &HelloServiceStub{conn}

}

func (c *HelloServiceStub) Hello(p *global.Person, res *global.Res) error {
	err := c.Call(handler.HelloServiceName+".Hello", p, res)
	if err != nil {
		return err
	}
	return nil
}
