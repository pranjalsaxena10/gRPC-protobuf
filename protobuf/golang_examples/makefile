SRC_DIR=$(PWD)/src/protos

build:
	protoc -I=$(SRC_DIR) --go_out=paths=source_relative:./generated $(SRC_DIR)/*.proto