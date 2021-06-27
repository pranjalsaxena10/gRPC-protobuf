package main

import (
	"io/ioutil"
	"log"
	"github.com/spf13/pflag"

	"github.com/pranjalsaxena10/gRPC-protobuf/protobuf/golang_examples/generated"
	"google.golang.org/protobuf/proto"
)


func main() {
	var fileName string
	pflag.StringVarP(&fileName, "fileName", "", "file.bin", "Name of file")
	pflag.Parse()
	SimpleMessageReadWriteFromFile(fileName)

}

func SimpleMessageReadWriteFromFile(fileName string) {
	simpleMessage := createSimpleMessageStruct()
	writeToFile(fileName, simpleMessage)
	readFromFile(fileName, simpleMessage)
}

func readFromFile(fileName string, simpleMessage *simplepb.SimpleMessage) {
	log.Print("readFromFile started...")
	
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error occurred while reading from ", fileName, "err: ", err)
	}

	err = proto.Unmarshal(data, simpleMessage)
	if err != nil {
		log.Fatal("Error occurred while unmarshalling ", simpleMessage, "err: ", err)
	}

	log.Printf("Reading from file is successful. Proto: %+v", simpleMessage)
}

func writeToFile(fileName string, simpleMessage *simplepb.SimpleMessage) {
	log.Print("writeToFile started...")

	protoData, err := proto.Marshal(simpleMessage)

	if err != nil {
		log.Fatal("Error occurred while marshalling ", simpleMessage, "err: ", err)
	}

	err = ioutil.WriteFile(fileName, protoData, 0644)

	if err != nil {
		log.Fatal("Error occurred while writing to file: ", fileName, "err : ", err)
	}

	log.Print("Writing to file is successful")
}



func createSimpleMessageStruct() *simplepb.SimpleMessage{
	value := simplepb.SimpleMessage{
		Id:          456,
		IsPresent:   true,
		MessageName: "message from proto",
		List:        []int32{5, 7, 8},
	}
	return &value
}
