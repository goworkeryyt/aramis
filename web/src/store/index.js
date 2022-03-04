import Vue from 'vue'
import Vuex from 'vuex'
import app from './modules/app'
import tagsView from './modules/tagsView'
import settings from './modules/settings'
import { user } from './modules/user'
import { router } from './modules/router'
import getters from './getters'
import VuexPersistence from 'vuex-persist'

Vue.use(Vuex)
const vuexSession = new VuexPersistence({
  storage: window.sessionStorage,
  modules: ['router', 'user']
})
const store = new Vuex.Store({
  modules: {
    app,
    tagsView,
    settings,
    user,
    router
  },
  getters,
  plugins: [vuexSession.plugin]
})

export default store
