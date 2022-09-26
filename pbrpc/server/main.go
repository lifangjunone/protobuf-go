package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"protobuf-go/pbrpc/service"
)

var _ service.HelloService = (*HelloService)(nil)

type HelloService struct{}

func (h *HelloService) Hello(request *service.Request, response *service.Response) error {
	response.Value = fmt.Sprintf("hello %s", request.Value)
	return nil
}

func main() {
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
