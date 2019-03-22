ENV:=dev
COMMAND:=
PROJECT_ID:=
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH:=
GBOOKSHELF_BOOKSHELF:=mybookshelf
GBOOKSHELF_SERVER_PORT:=2109
GBOOKSHELF_METRICS_PORT:=2112
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

run:
	docker run \
	-p $(GBOOKSHELF_SERVER_PORT):$(GBOOKSHELF_SERVER_PORT) \
	-p $(GBOOKSHELF_METRICS_PORT):$(GBOOKSHELF_METRICS_PORT) \
	-e "GBOOKSHELF_BOOKSHELF=${GBOOKSHELF_BOOKSHELF}" \
	-e "PROJECT_ID=$(PROJECT_ID)" \
	-e "GBOOKSHELF_SERVER_PORT=$(GBOOKSHELF_SERVER_PORT)" \
	-e "GBOOKSHELF_METRICS_PORT=$(GBOOKSHELF_METRICS_PORT)" \
	-e "FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=/credentials/firestore-adminsdk.json" \
	--mount type=bind,source=$(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH),target=/credentials/firestore-adminsdk.json,readonly \
	gbookshelf-server:local 

run-vue:
	cd web/gbookshelf-vue; yarn run dev

build-local:
	cloud-build-local --config=cloudbuild.yaml --dryrun=$(BUILD_DRYRUN) .

submit:
	gcloud builds submit --config cloudbuild.yaml .

run-gcp:
	gcloud auth configure-docker
	docker run \
	-p $(GBOOKSHELF_SERVER_PORT):$(GBOOKSHELF_SERVER_PORT) \
	-p $(GBOOKSHELF_METRICS_PORT):$(GBOOKSHELF_METRICS_PORT) \
	-e "GBOOKSHELF_BOOKSHELF=${GBOOKSHELF_BOOKSHELF}" \
	-e "GBOOKSHELF_SERVER_PORT=$(GBOOKSHELF_SERVER_PORT)" \
	-e "GBOOKSHELF_METRICS_PORT=$(GBOOKSHELF_METRICS_PORT)" \
	-e "PROJECT_ID=$(PROJECT_ID)" \
	-e "FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=/credentials/firestore-adminsdk.json" \
	--mount type=bind,source=$(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH),target=/credentials/firestore-adminsdk.json,readonly \
    gcr.io/$(PROJECT_ID)/gbookshelf-server:latest

drmi:
	docker rmi -f $$(docker images --filter "dangling=true" -q --no-trunc)

tf-apply:
	cd deployments/tf/; terraform apply

tf-destroy:
	cd deployments/tf/; terraform destroy

kube-init:
	gcloud container clusters get-credentials gbookshelf-dev --region asia-northeast1

kube-describles: kube-init
	kubectl get pods,deployments,daemonsets,services,endpoints,configmaps,persistentvolumeclaim,storageclass,namespaces,serviceaccount --show-labels --namespace $(ENV)-gbookshelf

# NOTE: Use envsubst until kustomize allows me to patch literal ConfigMap (https://github.com/kubernetes-sigs/kustomize/issues/680).

kube-build:
	export GBOOKSHELF_SERVICE=$(ENV)-gbookshelf-server; \
	export PROMETHUES_SERVICE=$(ENV)-prometheus; \
	kustomize build deployments/overlays/$(ENV) \
	| envsubst

kube-apply:
	export GBOOKSHELF_SERVICE=$(ENV)-gbookshelf-server; \
	export PROMETHUES_SERVICE=$(ENV)-prometheus; \
	kustomize build deployments/overlays/$(ENV) \
	| envsubst | kubectl apply -f -

kube-delete:
	export GBOOKSHELF_SERVICE=$(ENV)-gbookshelf-server; \
	export PROMETHUES_SERVICE=$(ENV)-prometheus; \
	kustomize build deployments/overlays/$(ENV) \
	| envsubst | kubectl delete -f -
