<template>
  <div id='bookshelf'>
    <AddBook />
    <Books />
  </div>
</template>

<script>
import Books from '../components/Books.vue'
import AddBook from '../components/AddBook.vue'

export default {
  name: 'Bookshelf',
  components: {
    Books,
    AddBook
  }
}
</script>

<style scoped>
#bookshelf {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 30px;
}
</style>