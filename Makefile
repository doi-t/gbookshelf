.PHONY: generate install

generate:
	protoc --go_out=. gbookshelf/gbookshelf.proto

install:
	go install ./cmd/gbookshelf
