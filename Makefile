.PHONY: generate install

generate:
	protoc --proto_path=api/protobuf-spec gbookshelf.proto --go_out=plugins=grpc:./pkg/apis/gbookshelf

install:
	go build -o server ./cmd/gbookshelf-server
	go install ./cmd/gbookshelf-client
