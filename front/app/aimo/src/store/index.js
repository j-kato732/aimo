import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    userId: 0
  },
  mutations: {
    setUserId (state, userId) {
    state.userId = userId
    }
  }
})