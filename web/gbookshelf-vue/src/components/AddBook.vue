<template>
  <div>
    <div class='add'>
      <form @submit='onSubmit'>
        <input type='text' v-model='title' placeholder='New Book Title...'>
        <input type='text' v-model='page' placeholder='New Book Page...'>
        <input type='text' v-model='why' placeholder='Why did you buy this book? (comming soon...)'>
        <input type='text' v-model='reason' placeholder='What do you expect from this book? (comming soon...)'>
        <input type='submit' value='Submit'>
      </form>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
export default {
  name: 'AddBook',
  data() {
    return {
      title: '',
      page: '',
      why: '',
      reason: ''
    }
  },
  methods: {
    ...mapActions(['addBook']),
    onSubmit(e) {
      e.preventDefault()
      const newBook = {
        title: this.title,
        page: this.page,
        current: 0, 
        done: false,
        why: this.why,
        reason: this.reason
      }
      this.addBook(newBook)

      // Clear submitted data
      this.title = ''
      this.title = ''
      this.page = ''
      this.why = ''
      this.reason = ''
    }
  }
}
</script>

<style scoped>
  form {
    display: flex;
    padding: 10px;
    margin-bottom: 10px;
  }
  input[type='text'] {
    flex: 10;
    padding: 10px;
    border: 1px solid #A0B084;
    outline: 0;
  }
  input[type='submit'] {
    flex: 2;
    color: #fff;
    background: #A0B084;
    border: 1px #A0B084 solid;
    cursor: pointer;
  }
</style>
