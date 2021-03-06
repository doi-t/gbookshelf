# init gbookshelf-vue

```sh
$ cd web/
$ yarn global add @vue/cli @vue/cli-service-global firebase-tools
$ firebase login
$ vue init webpack gbookshelf-vue
$ cd gbookshelf-vue/
$ yarn add grpc google-protobuf grpc-web
$ yarn add webpack-cli mini-css-extract-plugin
$ yarn upgrade --latest
$ yarn add babel-loader@7  # FIXME: Can't run 'yarn run dev' with babel-loader@8
$ yarn upgrade weback-dev-server@3.6.0 # FIXME: 3.7.0/3.7.1 causes "TypeError: Cannot assign to read only property 'exports' of object '#<Object>'"
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
```

## Notes of webpack complications
- https://vue-loader.vuejs.org/guide/#manual-setup
    - Update `build/webpack.base.conf.js`
- https://github.com/vuejs-templates/webpack/issues/1421#issuecomment-471209683
    - `vue init` generates webpack config for webpack 3
- https://github.com/webpack-contrib/uglifyjs-webpack-plugin/issues/362
    - Need to use terser-webpack-plugin to uglify ES6 code

# Deploy/Disable to Firebase Hosting
```
$ cd gbookshelf-vue/
$ yarn run build
$ firebase list # Check which project you are using
$ firebase use <Project ID> # Switch project
$ firebase deploy
$ # Setup TXT and A records for custom domain via Firebase UI
$ firebase hosting:disable
```

# Upgrade packages

```
$ cd gbookshelf-vue/
$ yarn list
$ yarn outdated
$ yarn upgrade-interactive --latest
```

# Upgrade global packages

```
$ cd gbookshelf-vue/
$ yarn global list
$ yarn global upgrade
```

# Test on local

```
$ echo 'GBOOKSHELF_SERVER_URL=http://<your domain>:8080'
$ make run-vue
$ open http://localhost:8080 # port number could be different depending on other opening ports.
```
