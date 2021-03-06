# gbookshelf
Manage your bookshelf to tackle a problem of many piled books which you would never read a.k.a. [Tundoku](https://en.wikipedia.org/wiki/Tsundoku).

[![Go Report Card](https://goreportcard.com/badge/github.com/doi-t/gbookshelf?style=flat-square)](https://goreportcard.com/report/github.com/doi-t/gbookshelf)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/doi-t/gbookshelf)
[![Release](https://img.shields.io/github/release/doi-t/gbookshelf.svg?style=flat-square)](https://github.com/doi-t/gbookshelf/releases/latest)

# Usage
**Super WIP: Currently it is pretty much useless.**

## Install Commands
```shell
make install
```

## Run Bookshelf Server
```shell
export GBOOKSHELF_BOOKSHELF=<Your bookshelf name (= Root Collection Name in Firestore Database)>; \
export GBOOKSHELF_SERVER_PORT=2109; \
export GBOOKSHELF_METRICS_PORT=2112; \
export PROJECT_ID=<Project ID>; \
export GCLOUD_CRENTIAL_FILE_PATH=$(pwd)/deployments/base/.credentials/gbookshelf-firebase-adminsdk.json; \
gbookshelf-server # in a terminal
```

## Access bookshelf
See help generated by [cobra](https://github.com/spf13/cobra).
```shell
gbsctl # in another terminal
gbsctl help [command]
```

## Test
```shell
make \
PROJECT_ID=<Project ID> \
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=$(pwd)/deployments/base/.credentials/gbookshelf-firebase-adminsdk.json \
test
```

## Build & Run
```shell
make build
make \
GBOOKSHELF_BOOKSHELF=<Your bookshelf name (= Root Collection Name in Firestore Database)> \
PROJECT_ID=<Project ID> \
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=</path/to/credentials> \
run
```

### Local build with cloud-build-local
```shell
make build-local
```

### Submit container image to GCR
```shell
make submit
```

### Run submitted container image on local
```shell
make submit
make \
GBOOKSHELF_BOOKSHELF=<Your bookshelf name (= Root Collection Name in Firestore Database)> \
PROJECT_ID=<Project ID> \
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=$(pwd)/deployments/base/.credentials/gbookshelf-firebase-adminsdk.json \
run-gcp
```

## Deploy

```shell
make tf-apply \
&& make \
ALERTMANAGER_SLACK_CHANNEL=<your channel name including '#'> \
ALERTMANAGER_SLACK_WEB_HOOK=<your slack incoming webhook url> \
kube-apply
```

### Access Services

```shell
GBOOKSHELF_ENV=dev; sudo kubefwd services --namespace ${GBOOKSHELF_ENV}-gbookshelf
```

```shell
GBOOKSHELF_ENV=dev; kubectl port-forward $(kubectl get pods --namespace ${GBOOKSHELF_ENV}-gbookshelf -l "name=gbookshelf-server" -o jsonpath="{.items[0].metadata.name}") 8080:8080 --namespace ${GBOOKSHELF_ENV}-gbookshelf
gbsctl list
```

## Destroy

```shell
make kube-delete
make tf-destroy
```

# References
- https://github.com/campoy/justforfunc
- https://github.com/golang-standards/project-layout
- https://godoc.org/github.com/golang/protobuf/proto
- https://godoc.org/cloud.google.com/go/firestore
- https://github.com/grpc/grpc-web
- https://www.envoyproxy.io/docs/envoy/latest/intro/version_history
