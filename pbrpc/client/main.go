package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"protobuf-go/pbrpc/service"
)

var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	client *rpc.Client
}

func NewHelloService(network, addr string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

func (h *HelloServiceClient) Hello(request *service.Request, response *service.Response) error {
	err := h.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	helloSvc, err := NewHelloService("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	var rsp service.Response
	err = helloSvc.Hello(&service.Request{
		Value: "ldd",
	}, &rsp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp.Value)
}
