package pb_test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"protobuf-go/pb"
	"testing"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)
	// object > protobuf > []byte
	msg := pb.String{
		Value: "hello",
	}
	bytes, err := proto.Marshal(&msg)
	if should.NoError(err) {
		fmt.Println(bytes)
	}
	// []byte > protobuf > object
	var mess = &pb.String{}
	err = proto.Unmarshal(bytes, mess)
	if should.NoError(err) {
		fmt.Println(mess.Value)
	}
}
