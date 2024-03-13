package main

import (
	"fmt"
	"go_rpc_demo/global"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:6666")
	if err != nil {
		panic("rpc Dial failed")
	}

	defer client.Close()

	var res global.Res
	var p = global.Person{
		Name: "dcl",
		Age:  21,
	}
	err = client.Call("HelloService.Hello", &p, &res)
	if err != nil {
		panic("调用失败" + err.Error())
	}
	fmt.Println(res)

}
