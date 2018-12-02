COMMAND:=
PROJECT_ID:=
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH:=
BOOKSHELF:=bookShelfTest

.PHONY: generate ensure install test add build run

generate:
	protoc --proto_path=api/protobuf-spec gbookshelf.proto --go_out=plugins=grpc:./pkg/apis/gbookshelf

ensure:
	dep ensure

install: ensure generate
	go install ./cmd/gbookshelf-server
	go install ./cmd/gbsctl

test: install
	./scripts/integration_test.sh $(PROJECT_ID) $(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH) $(BOOKSHELF)

add:
	./scripts/cobra_add.sh $(COMMAND)

build:
	docker build -f Dockerfile  -t gbookshelf-server:local .

run:
	docker run -p 8888:8888 -p 2112:2112 \
	-e "BOOKSHELF=${BOOKSHELF}" \
	-e "PROJECT_ID=$(PROJECT_ID)" \
	-e "FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=/credentials/firestore-adminsdk.json" \
	--mount type=bind,source=$(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH),target=/credentials/firestore-adminsdk.json,readonly \
	gbookshelf-server:local 
