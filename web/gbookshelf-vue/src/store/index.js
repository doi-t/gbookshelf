import Vuex from 'vuex'
import Vue from 'vue'
import books from './modules/books'

// Load Vuex
Vue.use(Vuex)

// Create store
export default new Vuex.Store({
  modules: {
    books
  }
})
