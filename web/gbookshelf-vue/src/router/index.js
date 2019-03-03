import Vue from 'vue'
import Router from 'vue-router'
import gbookshelf from '@/components/gbookshelf'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'gbookshelf',
      component: gbookshelf
    }
  ]
})
