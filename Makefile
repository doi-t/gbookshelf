.PHONY: generate install

generate:
	protoc gbookshelf/gbookshelf.proto --go_out=plugins=grpc:.

install:
	go build -o server ./cmd/server
	go install ./cmd/gbookshelf
