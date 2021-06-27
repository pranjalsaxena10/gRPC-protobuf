# Protobuf:
This is a sample project related to GoLang working of protobuf where:
1. You can add your Google Protobuf files
2. Create GoLang struct for protobuf files
3. Serialize them in binary files
4. Read binary files
5. Convert them back to Protobuf Message

-----

How to use this project:
go run main.go --protobufName name_of_your_protomessage --file name_of_binary_file

-----

Example: 
go run main.go --protobufName EnumMessage --file EnumMessage.bin

-----

Steps to add new protofiles to this project:
1. Create your protobuf file like simple.proto following google protobuf syntax
2. Put it inside directory src/protos
3. Run command "make build". This will compile newly added proto files inside dir "generated"
4. Add corresponding struct creation code like:
    This is for a proto file like:
    ```
        syntax = "proto3";
        option go_package = "github.com/pranjalsaxena10/gRPC-protobuf/protobuf/golang_examples;protob";

        message SimpleMessage {
            int32 id = 1;
            bool isPresent = 2;
            string messageName = 3;
            repeated int32 list = 4;
        }
    ```
    Struct creation:
    ```
        func createSimpleMessageStruct() *protob.SimpleMessage{
            simpleMessageProto := protob.SimpleMessage{
                Id:          456,
                IsPresent:   true,
                MessageName: "message from proto",
                List:        []int32{5, 7, 8},
            }
            return &simpleMessageProto
        }
    ```

5. Add corresponding update in createStruct(), readFromFile() and writeFromFile() methods for switch case
6. Once all this is done you can happily use
    go run main.go --protobufName name_of_your_protomessage --file name_of_binary_file


So, with minimal code changes you can add your proto file in this project, serialize them and store them in binary files.
Enjoy☺️!

