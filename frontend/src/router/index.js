import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '@/components/Login'
import Admin from '@/components/Admin'
Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'index',
      component: Index
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
    // {
    //   path: '/admin',
    //   name: 'admin',
    //   component: Admin,
    //   children: [
    //     {
    //       path: 'user',
    //       component: User,
    //       name: 'User'
    //     },
    //     {
    //       path: 'article'
    //     }
    //   ]
    // }
  ]
})
