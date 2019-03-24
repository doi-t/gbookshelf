<template>
  <div id='bookshelf'>
    <section>
      <span class='title-text'>No more Tsundoku</span>
        <div>
        <input v-model='inputField' v-on:keyup.enter='Add' placeholder='Your new book title'>
        <button @click='Add'>Add a new book to your bookshelf</button>
      </div>
    </section>
    <section>
      <div class='booklist'>
        <ul>
          <li
            v-for='book in books'
            v-bind:key='book.title'
          >
            <div>
              '{{book.title}}'
              {{book.page}}p
              ({{book.current}})
              done? {{book.done}}
              <button @click='Update(book)'>Update</button>
              <button @click='Remove(book)'>Remove</button>
            </div>
          </li>
        </ul>
      </div>
    </section>
  </div>
</template>

<script>
import { Void, Book, Books } from 'gbookshelf_pb'
import { BookShelfClient } from 'gbookshelf_grpc_web_pb'

export default {
  name: 'gbookshelf',
  components: {},
  data: function () {
    return {
      inputField: '', // new book title
      books: []
    }
  },
  created: function () {
    this.client = new BookShelfClient('http://dev-gbookshelf-server:8080', null, null) // TODO: make 'dev-' dynamic
    this.List()
  },
  methods: {
    List: function () {
      let voidRequest = new Void()
      this.client.list(voidRequest, {}, (err, response) => {
        this.books = response.toObject().booksList
        console.log(this.books)
      })
    },
    Add: function () {
      let bookRequest = new Book()
      bookRequest.setTitle(this.inputField)
      this.client.add(bookRequest, {}, (err, response) => {
        this.title= response.toObject().title
        console.log(this.title, 'has been added.')

        // refresh the page
        this.inputField = ''
        this.List()
      })
    },
    Update: function (book) {
      let bookRequest = new Book()
      bookRequest.setTitle(book.title)
      bookRequest.setPage(book.page)
      bookRequest.setCurrent(book.current)
      bookRequest.setDone(true)
      this.client.update(bookRequest, {}, (err, response) => {
        this.title= response.toObject().title
        console.log(this.title, 'has been updated.')

        // refresh the page
        this.inputField = ''
        this.List()
      })
    },
    Remove: function (book) {
      let bookRequest = new Book()
      bookRequest.setTitle(book.title)
      this.client.remove(bookRequest, {}, (err, response) => {
        this.title = response.toObject().title
        console.log(this.title, 'has been deleted.')

        // refresh the page
        this.inputField = ''
        this.List()
      })
    }
  }
}
</script>

<style>
#bookshelf {
  margin: 10px auto;
  max-width: 700px;
  line-height: 1.4;
}
</style>
