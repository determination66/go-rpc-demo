package main

import (
	"fmt"
	"go_rpc_demo/global"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		panic("net.Dial:" + err.Error())
	}
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var res global.Res

	var p = global.Person{
		Name: "张三",
		Age:  22,
	}

	err = client.Call("HelloService.Hello", &p, &res)
	if err != nil {
		panic("conn.Call:" + err.Error())
	}

	fmt.Println(res)

}
