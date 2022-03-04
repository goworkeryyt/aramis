
const state = {
  visitedViews: [],
  activePath: '/',
  activeName: ''
}

const mutations = {
  ADD_VISITED_VIEW: (state, view) => {
    if (state.visitedViews.some(v => v.path === view.path)) return

    state.visitedViews.push(view)
  },
  DEL_VISITED_VIEW: (state, view) => {
    for (const [i, v] of state.visitedViews.entries()) {
      if (v.path === view) {
        state.visitedViews.splice(i, 1)

        break
      }
    }
  },
  DEL_ALL_VISITED_VIEWS: state => {
    state.visitedViews = []
  },
  DEL_OTHER_VISITED_VIEWS: (state, length) => {
    state.visitedViews.splice(0, length)
  },
  SET_ACTIVE_PATH: (state, path) => {
    state.activePath = path
  }
}

const actions = {
  addView({ dispatch }, view) {
    dispatch('addVisitedView', view)
  },
  addVisitedView({ commit }, view) {
    commit('ADD_VISITED_VIEW', view)
  },
  delView({ dispatch, state }, view) {
    return new Promise(resolve => {
      dispatch('delVisitedView', view)
      resolve({
        visitedViews: [...state.visitedViews]
      })
    })
  },
  delVisitedView({ commit, state }, view) {
    return new Promise(resolve => {
      commit('DEL_VISITED_VIEW', view)
      resolve([...state.visitedViews])
    })
  },
  delAllVisitedViews({ commit, state }) {
    return new Promise(resolve => {
      commit('DEL_ALL_VISITED_VIEWS')
      resolve([...state.visitedViews])
    })
  },
  delOtherVisitedViews({ commit, state }, length) {
    return new Promise(resolve => {
      commit('DEL_OTHER_VISITED_VIEWS', length)
      resolve([...state.visitedViews])
    })
  },
  setActivePath({ commit, state }, path) {
    commit('SET_ACTIVE_PATH', path)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
