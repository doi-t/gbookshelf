## init gbookshelf-vue

```sh
$ cd web/
$ yarn global add @vue/cli firebase-tools
$ firebase login
$ vue init webpack gbookshelf-vue
$ cd gbookshelf-vue
$ yarn add grpc google-protobuf grpc-web webpack-cli babel-loader@7
$ yarn upgrade --latest
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
$ firebase deploy
$ # Setup TXT and A records for custom domain via Firebase UI
```

### Upgrade packages

```
$ yarn list
$ yarn outdated
$ yarn upgrade-interactive --latest
```

### Upgrade global packages

```
$ yarn global list
$ yarn global upgrade
```

## Test on local

```
$ make run-vue
$ open http://localhost:8081
```
