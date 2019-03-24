<template>
  <div class='book-item'>
    <div class='book-container'>
      <div class='book-title'>
        {{ book.title }}
      </div>
      <div class='book-page'>
        {{ book.current }} / {{ book.page }} pages ({{ bookProgress(book) }}%)
      </div>
      <div class='book-status'>
        {{ book.done }}
      </div>
      <div class='book-update'>
        <form @submit='onSubmit(book, current)'>
          <input type='text' v-model='current' placeholder='Input Current Page...'>
          <input type='submit' value='Update'>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default { name: 'BookItem',
  props: ['book'],
  data() {
    return {
      current: ''
    }
  },
  methods: {
    ...mapActions(['updateBook']),
    bookProgress: function (book) {
      return Math.round(100 * book.current / book.page * 100) / 100
    },
    onSubmit(book, newCurrent, e) {
      if (e) e.preventDefault()
      const updBook = {
        title: book.title,
        page: book.page,
        current: newCurrent, 
        done: book.done
      }
      this.updateBook(updBook)
      this.current = ''
    }
  }
}
</script>

<style scoped>
  .book-container {
    display: grid;
  }
  .book-title {
    grid-column: 1 / span 2;
    grid-row: 1;
  }
  .book-page {
    grid-column: 1;
    grid-row: 2;
  }
  .book-status {
    grid-column: 2;
    grid-row: 2;
  }
  .book-update {
    grid-column: 1;
    grid-row: 3;
  }
  form {
    display: flex;
    padding: 10px;
  }
  input[type='text'] {
    flex: 10;
    padding: 10px;
    border: 1px solid #A57C65;
    outline: 0;
  }
  input[type='submit'] {
    flex: 2;
    color: #fff;
    background: #A57C65;
    border: 1px #A57C65 solid;
    cursor: pointer;
  }
</style>
