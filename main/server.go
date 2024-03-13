package main

import (
	"fmt"
	"go_rpc_demo/global"
	"io"
	"net/http"
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

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			panic("rpc.ServeRequest" + err.Error())
		}
	})

	_ = http.ListenAndServe(":6666", nil)

}
