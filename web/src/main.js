import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Element from 'element-ui'
import 'babel-polyfill'
import 'normalize.css/normalize.css'
import '@/styles/index.scss'
import '@/styles/element-variables.scss'
import '@/styles/variables.scss'
import '@/styles/btn.scss'

import constant from './commen/constant/index.js'
import formatter from './commen/constant/formatter.js'
import globalComponents from './utils/globalComponents.js'
import globalFunction from './utils/globalFunction.js'
import api from './http/index.js'

Vue.use(Element, {
  size: 'medium' // set element-ui default size
})

Vue.use(globalComponents)

Vue.config.productionTip = false

Vue.prototype.Constant = constant
Vue.prototype.Formatter = formatter
Vue.prototype.GlobalFunction = globalFunction
Vue.prototype.Api = api

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
