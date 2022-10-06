import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: () =>
        import ('../components/HelloWorld.vue')
    },
    {
      path: '/recommend',
      name: 'Recommend',
      component: () =>
        import ('../components/Recommend.vue')
    }
  ]
})
