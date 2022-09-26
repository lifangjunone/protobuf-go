## protoc generate proto.go file
### command
```shell

protoc -I=. --go_out=./ --go_opt=module="protobuf-go/pb" ./hello.proto
```