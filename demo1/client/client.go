package main

import (
	"context"
	"fmt"
	"go_rpc_demo/protoc/hello_world"
	"google.golang.org/grpc"
	"time"
)

const (
	ADDRESS = "localhost:50052"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := hello_world.NewGreeterClient(conn)
	//调用服务端推送流
	//reqstreamData := &stream.StreamReqData{Data: "aaa"}
	helloReq := &hello_world.HelloRequest{
		Name:    "张三",
		Age:     16,
		Courses: []string{"gin", "go", "protobuf"},
		Sex:     hello_world.Sex_Male,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	res, err := c.SayHello(ctx, helloReq)
	if err != nil {
		panic("err:" + err.Error())
	}

	fmt.Println(res)
	cancel()

}
