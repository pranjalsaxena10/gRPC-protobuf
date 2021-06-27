SRC_DIR=$PWD/src/protos
protoc -I=$SRC_DIR --go_out=paths=source_relative:./generated $SRC_DIR/simple.proto