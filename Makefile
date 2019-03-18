COMMAND:=
PROJECT_ID:=
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH:=
BOOKSHELF:=bookShelfTest
BUILD_DRYRUN:=false

.PHONY: generate ensure install test add build run

generate:
	@protoc --proto_path=api/protobuf-spec gbookshelf.proto \
		--go_out=plugins=grpc:./pkg/apis/gbookshelf/
	@mkdir -p ./web/gbookshelf-vue/node_modules
	@protoc --proto_path=api/protobuf-spec gbookshelf.proto \
		--js_out=import_style=commonjs,binary:./web/gbookshelf-vue/node_modules/ \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web/gbookshelf-vue/node_modules/
	 
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
	docker build -f Dockerfile -t gbookshelf-server:local .
	docker build -f build/Dockerfile.envoy -t envoy:local .

run:
	docker run -p 8888:8888 -p 2112:2112 \
	-e "BOOKSHELF=${BOOKSHELF}" \
	-e "PROJECT_ID=$(PROJECT_ID)" \
	-e "FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=/credentials/firestore-adminsdk.json" \
	--mount type=bind,source=$(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH),target=/credentials/firestore-adminsdk.json,readonly \
	gbookshelf-server:local 

run-envoy:
	docker run -p 8080:8080 envoy:local

run-vue:
	cd web/gbookshelf-vue
	yarn run dev

build-local:
	cloud-build-local --config=cloudbuild.yaml --dryrun=$(BUILD_DRYRUN) .

submit:
	gcloud builds submit --config cloudbuild.yaml .

run-gcp:
	gcloud auth configure-docker
	docker run -p 8888:8888 -p 2112:2112 \
	-e "BOOKSHELF=${BOOKSHELF}" \
	-e "PROJECT_ID=$(PROJECT_ID)" \
	-e "FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=/credentials/firestore-adminsdk.json" \
	--mount type=bind,source=$(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH),target=/credentials/firestore-adminsdk.json,readonly \
    gcr.io/$(PROJECT_ID)/gbookshelf-server:latest

drmi:
	docker rmi $(docker images --filter "dangling=true" -q --no-trunc)

kube-describles:
	gcloud container clusters get-credentials gbookshelf-dev --region asia-northeast1
	kubectl get pods,deployments,daemonsets,services,endpoints,configmaps,persistentvolumeclaim,storageclass,namespaces,serviceaccount --show-labels --namespace gbookshelf-server

# TODO: update 'base' to overlay name accordingly
kube-build:
	kustomize build deployments/base

kube-apply:
	kustomize build deployments/base | kubectl apply -f -

kube-delete:
	kustomize build deployments/base | kubectl delete -f -
