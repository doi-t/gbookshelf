import Vue from 'vue'
import Router from 'vue-router'
import Bookshelf from '@/views/Bookshelf.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'bookshelf',
      component: Bookshelf
    }
  ]
})
