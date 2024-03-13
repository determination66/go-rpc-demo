package main

import (
	"fmt"
	"go_rpc_demo/global"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "localhost:6666")
	if err != nil {
		panic("net.Dial:" + err.Error())
	}
	defer conn.Close()

	var res global.Res

	var p = global.Person{
		Name: "张三",
		Age:  22,
	}

	err = conn.Call("HelloService.Hello", &p, &res)
	if err != nil {
		panic("conn.Call:" + err.Error())
	}

	fmt.Println(res)

}
