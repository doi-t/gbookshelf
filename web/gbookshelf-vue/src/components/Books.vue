<template>
  <div>
    <div class='books'>
      <!-- FIXME: key should be id in DB -->
      <div 
        @dblclick='onDblClick(book)' 
        v-for='book in allBooks' 
        v-bind:key='book.title' 
        class='book'
        v-bind:class='{"is-complete": book.done}'
      > 
      <BookItem v-bind:book='book' />
      <i @click='removeBook(book.title)' class='fas fa-trash-alt'></i>
      </div>
    </div>
  </div>
</template>

<script>
import BookItem from './BookItem.vue'
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'Books',
  components: {
    BookItem
  },
  methods: {
    ...mapActions(['fetchBooks', 'removeBook', 'updateBook']),
    onDblClick(book) {
      const updBook = {
        title: book.title,
        page: book.page,
        current: book.current,
        done: !book.done
      }
      this.updateBook(updBook)
    }
  },
  computed: {
    ...mapGetters(['allBooks'])
  },
  created() {
    this.fetchBooks()
  }
}
</script>

<style scoped>
  .books {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 1rem;
  }
  .book {
    border: 1px solid #ccc;
    color: #fff;
    background: #A0B084;
    padding: 1rem;
    border-radius: 5px;
    text-align: center;
    position: relative;
    cursor: pointer;
  }
  i {
    position: absolute;
    bottom: 10px;
    right: 10px;
    color: #fff;
    cursor: pointer;
  }
  .legend {
    display: flex;
    justify-content: space-around;
    margin-bottom: 1rem;
  }
  .complete-box {
    display: inline-block;
    width: 10px;
    height: 10px;
    background: #688B8A;
  }
  .incomplete-box {
    display: inline-block;
    width: 10px;
    height: 10px;
    color: #fff;
    background: #A0B084;
  }
  .is-complete {
    background: #688B8A;
    color: #fff;
  }
  @media (max-width: 500px) {
    .books {
      grid-template-columns: 1fr;
    }
  }
</style>
