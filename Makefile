.PHONY: generate install

generate:
	protoc --proto_path=api/protobuf-spec gbookshelf.proto --go_out=plugins=grpc:./pkg/apis/gbookshelf

install:
	go install ./cmd/gbookshelf-server
	go install ./cmd/gbsctl

test: install
	gbookshelf-server &
	rm -f mydb.pb
	gbsctl add hoge
	gbsctl add fuga
	gbsctl add foo
	gbsctl list
	ls -lh mydb.pb
	xxd mydb.pb
	gbsctl remove fuga
	ls -lh mydb.pb
	xxd mydb.pb
	gbsctl list
	pkill gbookshelf-server
