## init gbookshelf-vue

```sh
$ cd web/
$ yarn global add vue-cli firebase-tools
$ firebase login
$ vue init webpack gbookshelf-vue
$ cd gbookshelf-vue
$ yarn run build
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
$ yarn add grpc google-protobuf grpc-web webpack-cli @babel/core @babel/preset-env
$ firebase deploy
$ # Setup TXT and A records for custom domain via Firebase UI
```

## Test on local

```
$ make run-vue
$ open http://localhost:8081
```
