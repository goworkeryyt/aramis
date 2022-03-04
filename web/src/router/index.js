/*
 * @Description: 路由
 */
import Vue from 'vue'
import VueRouter from 'vue-router'

import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

import store from '../store'
import routes from './router'

Vue.use(VueRouter)

NProgress.inc(0.2)
NProgress.configure({ easing: 'ease', speed: 500, showSpinner: false })

const router = new VueRouter({
  mode: 'hash',
  routes
})

router.beforeEach((to, from, next) => {
  // console.log(sessionStorage.length);
  // if (to.path === "/") {
  //   sessionStorage.clear();
  //   // console.log(sessionStorage);
  //   // window.location.reload();
  // }
  next()
})

router.afterEach(() => {
  NProgress.done()
})

export default router
