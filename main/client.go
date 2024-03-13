package main

import (
	"fmt"
	"go_rpc_demo/client_proxy"
	"go_rpc_demo/global"
)

func main() {
	client := client_proxy.NewHelloServiceStub("tcp", "localhost:6666")

	var res global.Res
	var p = global.Person{
		Name: "dcl",
		Age:  21,
	}

	err := client.Hello(&p, &res)
	if err != nil {
		panic("client.Hello:" + err.Error())
	}

	fmt.Println(res)
}
