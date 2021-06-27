package main

import (
	"fmt"
	"github.com/pranjalsaxena10/gRPC-protobuf/protobuf/golang_examples/generated"
)

func main() {
	createSimpleMessageStruct()
}

func createSimpleMessageStruct() {
	value := simplepb.SimpleMessage{
		Id:          456,
		IsPresent:   true,
		MessageName: "message from proto",
		List:        []int32{5, 7, 8},
	}
	fmt.Println(value)

}
