import Authority from '../../http/Authority.js'
export const router = {
  namespaced: true,
  state: {
    userMenuList: []
  },
  mutations: {
    setUserMenuList(state, userMenuList) {
      state.userMenuList = userMenuList
    }
  },
  actions: {
    async getSelfMenuTree({ commit }) {
      const res = await Authority.getSelfMenuTree()
      if (res.data.code === '0') {
        commit('setUserMenuList', res.data.content)
      }
    }
  }
}
