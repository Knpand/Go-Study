// import Vue from 'vue'
// import VueRouter from "vue-router"


import { createRouter, createWebHashHistory } from 'vue-router'
import Login from './components/LoginView.vue'
import Top from './components/TopView.vue'
import { authorizeToken } from './guards'


const routes = [
  {
    path: '/',
    name: 'Top',
    component: Top,
    meta: {
      requiresAuth: true
    }

  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})
router.beforeEach(authorizeToken)

export default router