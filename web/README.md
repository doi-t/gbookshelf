## init gbookshelf-vue

```sh
$ cd web/
$ npm i vue-cli -g
$ npm install -g firebase-tools
$ firebase login
$ mkdir gbookshelf-vue
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
$ firebase deploy
$ # Setup TXT and A records for custom domain via Firebase UI
```

## Test on local
```
$ npm run dev
$ open http://localhost:8080
```
