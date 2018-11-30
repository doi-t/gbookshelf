COMMAND:=
PROJECT_ID:=
GCLOUD_CRENTIAL_FILE_PATH:=
BOOKSHELF:=bookShelfTest

.PHONY: generate, ensure, install, test, add

generate:
	protoc --proto_path=api/protobuf-spec gbookshelf.proto --go_out=plugins=grpc:./pkg/apis/gbookshelf

ensure:
	dep ensure

install: ensure generate
	go install ./cmd/gbookshelf-server
	go install ./cmd/gbsctl

test: install
	./scripts/integration_test.sh $(PROJECT_ID) $(GCLOUD_CRENTIAL_FILE_PATH) $(BOOKSHELF)



add:
	./scripts/cobra_add.sh $(COMMAND)
