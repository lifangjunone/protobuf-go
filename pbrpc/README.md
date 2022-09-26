## go code generate
```shell
protoc -I=. --go_out=./.. --go_opt=module="protobuf-go/pbrpc/service" hello.proto
```