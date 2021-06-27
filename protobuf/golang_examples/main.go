package main

import (
	"io/ioutil"
	"log"
	"github.com/pranjalsaxena10/gRPC-protobuf/protobuf/golang_examples/generated"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var protobufName, fileName string

func main() {
	InitializeFlag()
	Init(protobufName, fileName)
}

func InitializeFlag() {
	pflag.StringVarP(&protobufName, "protobufName", "", "SimpleMessage", "Name of Protobuf")
	pflag.StringVarP(&fileName, "file", "", "file.bin", "Binary file to store serialized proto")
	pflag.Parse()
}

func Init(protobufName string, fileName string) {
	ReadWriteFromFile(protobufName, fileName)
}

func ReadWriteFromFile(protobufName string, fileName string) {
	protobufStruct := createStruct(protobufName)
	writeToFile(protobufName, fileName, protobufStruct)
	readFromFile(protobufName, fileName, protobufStruct)
}

func createStruct(protobufName string) interface{} {
	switch protobufName {
	case "EnumMessage":
		return createEnumMessageStruct()
	default:
		return createSimpleMessageStruct()
	}
}

func readFromFile(protobufName string, fileName string, protobufStruct interface{}) {
	log.Printf("[%s][%s] readFromFile() started...", protobufName, fileName)
	
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("[%s] Error occurred while reading from %s, error as: %v", protobufName, fileName, err)
	}

	var m protoreflect.ProtoMessage
	switch protobufName {
	case "EnumMessage":
		m = protobufStruct.(*protob.EnumMessage)
	default:
		m = protobufStruct.(*protob.SimpleMessage)
	}

	err = proto.Unmarshal(data, m)

	if err != nil {
		log.Fatalf("[%s] Error occurred while unmarshalling %v, error as: %v", protobufName, m, err)
	}

	log.Printf("[%s] Reading from file %s is successful. Proto: %+v", protobufName, fileName, m)
}

func writeToFile(protobufName string, fileName string, protobufStruct interface{}) {
	log.Printf("[%s] writeToFile() started...", protobufName)

	var m protoreflect.ProtoMessage
	switch protobufName {
	case "EnumMessage":
		m = protobufStruct.(*protob.EnumMessage)
	default:
		m = protobufStruct.(*protob.SimpleMessage)
	}

	protoData, err := proto.Marshal(m)

	if err != nil {
		log.Fatalf("[%s] Error occurred while marshalling: %v error as: %v", protobufName, m, err)
	}

	err = ioutil.WriteFile(fileName, protoData, 0644)

	if err != nil {
		log.Fatalf("[%s] Error occurred while writing to file: %s error as: %v", protobufName, fileName, err)
	}

	log.Printf("[%s] Writing to file is successful", protobufName)
}

func createSimpleMessageStruct() *protob.SimpleMessage{
	simpleMessageProto := protob.SimpleMessage{
		Id:          456,
		IsPresent:   true,
		MessageName: "message from proto",
		List:        []int32{5, 7, 8},
	}
	return &simpleMessageProto
}

func createEnumMessageStruct() *protob.EnumMessage{
	enumMessageProto := protob.EnumMessage{
		Id:           78,
		DayOfTheWeek: protob.DayOfTheWeek_MONDAY,
	}
	return &enumMessageProto
}
