<template>
  <header class='header'>
    <div id='books-emoji'>📚 No More Tsundoku.</div> 
    <div id='nav'>
      <router-link to='/'>Bookshelf</router-link> |
      <router-link to='/about'>About</router-link>
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header'
}
</script>

<style scoped>
  .header {
    background: #A57C65;
    color: #fff;
    text-align: center;
    padding: 10px;
  }
  .header a {
    color: #fff;
    padding-right: 5px;
    text-decoration: none;
  }
  #books-emoji {
    font-size: 30pt
  }
</style>
