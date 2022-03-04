import Authority from '../../http/Authority.js'
export const user = {
  namespaced: true,
  state: {
    perms: [] // 用户权限标识集合
  },
  mutations: {
    setButtonPerms(state, perms) { // 用户权限标识集合
      state.perms = perms
    }
  },
  actions: {
    async getSelfButtons({ commit }) {
      const res = await Authority.getSelfButtons()

      if (res.data.code === '0') {
        commit('setButtonPerms', res.data.content)
      }
    }
  }
}
