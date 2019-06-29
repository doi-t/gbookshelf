ENV:=dev
COMMAND:=
PROJECT_ID:=
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH:=
ALERTMANAGER_SLACK_CHANNEL:=
ALERTMANAGER_SLACK_WEB_HOOK:=
GBOOKSHELF_DOMAIN:=
GBOOKSHELF_BOOKSHELF:=mybookshelf
GBOOKSHELF_SERVER_PORT:=2109
GBOOKSHELF_METRICS_PORT:=2112
CERT_MANAGER_LETS_ENCRYPT_EMAIL:=
CERT_MANAGER_ROUTE53_ACCESS_KEY_ID:=
BUILD_DRYRUN:=false

.PHONY: generate ensure install test add build run

generate:
	@protoc --proto_path=api/protobuf-spec gbookshelf.proto \
		--go_out=plugins=grpc:./pkg/apis/gbookshelf/
	@mkdir -p ./web/gbookshelf-vue/node_modules
	@protoc --proto_path=api/protobuf-spec gbookshelf.proto \
		--js_out=import_style=commonjs,binary:./web/gbookshelf-vue/node_modules/ \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web/gbookshelf-vue/node_modules/
	 
install: generate
	go install ./cmd/gbookshelf-server
	go install ./cmd/gbsctl

test: install
	./scripts/integration_test.sh $(PROJECT_ID) $(FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH) $(BOOKSHELF)

check-config:
	export PROMETHUES_SERVICE=dummy-prometheus; \
	export ALERTMANAGER_SERVICE=dummy-alertmanager; \
	cat deployments/base/prometheus/prometheus.yaml | envsubst > /tmp/prometheus.yaml; \
	promtool check config /tmp/prometheus.yaml; \
	cat deployments/base/prometheus/rules.yaml | envsubst > /tmp/rules.yaml; \
	promtool check rules /tmp/rules.yaml
	export ALERTMANAGER_SLACK_CHANNEL='#dummy'; \
	export ALERTMANAGER_SLACK_WEB_HOOK='https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX'; \
	cat ./deployments/base/alertmanager/alertmanager.yaml | envsubst | amtool check-config

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

run-vue: generate
	cd web/gbookshelf-vue; \
	yarn start 

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

open-port:
	gcloud compute firewall-rules create gbookshelf-firewall --allow tcp:30080

check-port:
	gcloud compute firewall-rules describe gbookshelf-firewall

close-port:
	gcloud compute firewall-rules delete gbookshelf-firewall

tf-apply:
	cd deployments/tf/; terraform apply

tf-destroy:
	cd deployments/tf/; terraform destroy

kube-init:
	gcloud container clusters get-credentials gbookshelf-dev --region asia-northeast1

kube-gets: kube-init
	kubectl get nodes --output wide --namespace $(ENV)-gbookshelf
	kubectl get pods,deployments,daemonsets,services,endpoints,ingresses,configmaps,secrets,persistentvolumeclaim,storageclass,namespaces,serviceaccount --show-labels --namespace $(ENV)-gbookshelf

# NOTE: Use envsubst until kustomize allows me to patch literal ConfigMap (https://github.com/kubernetes-sigs/kustomize/issues/680).

kube-build: check-config
	export GBOOKSHELF_SERVICE=$(ENV)-gbookshelf-server; \
	export GBOOKSHELF_DOMAIN=$(GBOOKSHELF_DOMAIN); \
	export PROMETHUES_SERVICE=$(ENV)-prometheus; \
	export ALERTMANAGER_SERVICE=$(ENV)-alertmanager; \
	export ALERTMANAGER_SLACK_CHANNEL=$(ALERTMANAGER_SLACK_CHANNEL); \
	export ALERTMANAGER_SLACK_WEB_HOOK=$(ALERTMANAGER_SLACK_WEB_HOOK); \
	kustomize build deployments/overlays/$(ENV) \
	| envsubst

kube-apply: kube-init check-config
	export GBOOKSHELF_SERVICE=$(ENV)-gbookshelf-server; \
	export GBOOKSHELF_DOMAIN=$(GBOOKSHELF_DOMAIN); \
	export PROMETHUES_SERVICE=$(ENV)-prometheus; \
	export ALERTMANAGER_SERVICE=$(ENV)-alertmanager; \
	export ALERTMANAGER_SLACK_CHANNEL=$(ALERTMANAGER_SLACK_CHANNEL); \
	export ALERTMANAGER_SLACK_WEB_HOOK=$(ALERTMANAGER_SLACK_WEB_HOOK); \
	kustomize build deployments/overlays/$(ENV) \
	| envsubst | kubectl apply -f -

kube-delete: kube-init
	export GBOOKSHELF_SERVICE=$(ENV)-gbookshelf-server; \
	export GBOOKSHELF_DOMAIN=$(GBOOKSHELF_DOMAIN); \
	export PROMETHUES_SERVICE=$(ENV)-prometheus; \
	export ALERTMANAGER_SERVICE=$(ENV)-alertmanager; \
	export ALERTMANAGER_SLACK_CHANNEL=$(ALERTMANAGER_SLACK_CHANNEL); \
	export ALERTMANAGER_SLACK_WEB_HOOK=$(ALERTMANAGER_SLACK_WEB_HOOK); \
	kustomize build deployments/overlays/$(ENV) \
	| envsubst | kubectl delete -f -

cert-manager-build:
	export GBOOKSHELF_DOMAIN=$(GBOOKSHELF_DOMAIN); \
	export CERT_MANAGER_LETS_ENCRYPT_EMAIL=$(CERT_MANAGER_LETS_ENCRYPT_EMAIL); \
	export CERT_MANAGER_ROUTE53_ACCESS_KEY_ID=$(CERT_MANAGER_ROUTE53_ACCESS_KEY_ID); \
	kustomize build deployments/cert-manager \
	| envsubst 

cert-manager-apply:
	export GBOOKSHELF_DOMAIN=$(GBOOKSHELF_DOMAIN); \
	export CERT_MANAGER_LETS_ENCRYPT_EMAIL=$(CERT_MANAGER_LETS_ENCRYPT_EMAIL); \
	export CERT_MANAGER_ROUTE53_ACCESS_KEY_ID=$(CERT_MANAGER_ROUTE53_ACCESS_KEY_ID); \
	kustomize build deployments/cert-manager \
	| envsubst | kubectl apply -f -

cert-manager-delete:
	export GBOOKSHELF_DOMAIN=$(GBOOKSHELF_DOMAIN); \
	export CERT_MANAGER_LETS_ENCRYPT_EMAIL=$(CERT_MANAGER_LETS_ENCRYPT_EMAIL); \
	export CERT_MANAGER_ROUTE53_ACCESS_KEY_ID=$(CERT_MANAGER_ROUTE53_ACCESS_KEY_ID); \
	kustomize build deployments/cert-manager \
	| envsubst | kubectl delete -f -

cert-manager-gets: kube-init
	kubectl get pods,deployments,daemonsets,services,endpoints,ingresses,configmaps,secrets,persistentvolumeclaim,storageclass,namespaces,serviceaccount,clusterissuer,certificate --show-labels --namespace cert-manager

cert-manager-check-certs: 
	kubectl get secrets stg-gbookshelf-tls --namespace cert-manager -o json | jq -r '.data."tls.crt"' | base64 -D | openssl x509 -text
	kubectl get secrets prd-gbookshelf-tls --namespace cert-manager -o json | jq -r '.data."tls.crt"' | base64 -D | openssl x509 -text
