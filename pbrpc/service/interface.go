package service

const (
	SERVICE_NAME = "HelloService"
)

type HelloService interface {
	Hello(request *Request, response *Response) error
}
