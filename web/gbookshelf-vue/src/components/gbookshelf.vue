<template>
  <div class='row'>
    <li
      v-for='book in books'
      v-bind:key='book.title'
    >
      <div>
        <span>{{book.title}}</span>
      </div>
    </li>
  </div>
</template>

<script>
import { Void } from 'gbookshelf_pb'
import { BookShelfClient } from 'gbookshelf_grpc_web_pb'

export default {
  name: 'gbookshelf',
  components: {},
  data: function () {
    return {
      inputField: '',
      books: []
    }
  },
  created: function () {
    this.client = new BookShelfClient('http://localhost:8080', null, null)
    this.List()
  },
  methods: {
    List: function () {
      let voidRequest = new Void()
      this.client.list(voidRequest, {}, (err, response) => {
        this.books = response.toObject().booksList
        console.log(this.books)
      })
    }
  }
}
</script>
