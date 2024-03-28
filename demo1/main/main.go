package main

import (
	"encoding/json"
	"fmt"
	"go_rpc_demo/demo1/protoc/hello_world"
	"google.golang.org/protobuf/proto"
)

func main() {
	req := hello_world.HelloRequest{
		Name:    "Bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	// proto 解析
	protoRsp, _ := proto.Marshal(&req)
	fmt.Println(len(protoRsp))

	// json解析
	jsonRsp, _ := json.Marshal(&req)
	fmt.Println(len(jsonRsp))

	//Unmarshal proto的二进制内容
	newReq := hello_world.HelloRequest{}
	_ = proto.Unmarshal(protoRsp, &newReq)
	fmt.Println(newReq.Courses)

}
