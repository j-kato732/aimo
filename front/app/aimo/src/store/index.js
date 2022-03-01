import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  getters:{
    userId: state => state.userId
  },
  state: {
    userId: 0
  },
  mutations: {
    setUserId (state, userId) {
    state.userId = userId
    }
  }
})