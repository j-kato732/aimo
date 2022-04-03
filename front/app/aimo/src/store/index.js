import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  getters:{
    userId: state => state.userId,
    period: state => state.period,
    period_management: state => state.period_m
  },
  state: {
    userId: 0,
    period: 0,
    period_management: 0,
  },
  mutations: {
    setUserId (state, userId) {
    state.userId = userId
    },
    setPeriod (state, period) {
      state.period = period
    },
    setPeriodForManagement (state, period_management) {
      state.period_management = period_management
    }
  }
})