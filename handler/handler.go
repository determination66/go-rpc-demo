package handler

import (
	"fmt"
	"go_rpc_demo/global"
)

// HelloServiceName 解决名字冲突
const HelloServiceName = "handler/HelloService"

type HelloService struct {
}

func (h *HelloService) Hello(p *global.Person, res *global.Res) error {
	res.Code = 200
	res.Msg = "hello," + p.Name
	fmt.Println(p)
	return nil
}
