## init gbookshelf-vue

```sh
$ cd web/
$ yarn global add vue-cli firebase-tools
$ firebase login
$ mkdir gbookshelf-vue
$ vue init webpack gbookshelf-vue
$ cd gbookshelf-vue $ yarn run build
$ firebase init hosting # ? What do you want to use as your public directory? dist
$ cat firebase.json
{
  "hosting": {
    "public": "dist",
    "ignore": [
      "firebase.json",
      "**/.*",
      "**/node_modules/**"
    ]
  }
}
$ firebase deploy
$ # Setup TXT and A records for custom domain via Firebase UI
```

## Test on local
```
$ make build
$ make \
BOOKSHELF=<Your bookshelf name (= Root Collection Name in Firestore Database)> \
PROJECT_ID=<Project ID> \
FIRESTORE_ADMINSDK_CRENTIAL_FILE_PATH=</path/to/credentials> \
run
```

```
$ make build
$ make run-envoy
```

```
$ cd web/gbookshelf-vue
$ yarn add grpc google-protobuf grpc-web
$ yarn run dev
$ open http://localhost:8081
```

```
$ docker kill $(docker ps | grep envoy:local | awk '{ print $1 }')
```
