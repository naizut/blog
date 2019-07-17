import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Login from '@/components/Login'
import Admin from '@/components/admin/Admin'
import UserAdmin from '@/components/admin/User'
import ArticleAdmin from '@/components/admin/Article'

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
    },
    {
      path: '/admin',
      name: 'admin',
      component: Admin
    },
    {
      path: '/admin/user',
      name: 'useradmin',
      component: UserAdmin
    },
    {
      path: '/admin/article',
      name: 'articleadmin',
      component: ArticleAdmin
    }
  ]
})
