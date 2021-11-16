import Vue from 'vue'
import VueRouter from 'vue-router'
import MyPage from '../views/MyPage.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/mypage',
    name: 'MyPage',
    component: MyPage
  },
  {
    path: '/',
    name: 'HelloWorld',
    component: () => import('../components/HelloWorld.vue')
  },
  {
    path: '/setting_user',
    name: 'Setting_user',
    component: () => import('../views/Setting_user.vue')
  },
  {
    path: '/aimSettingSheet',
    name: 'AimSettingSheet',
    component: () => import('../views/AimSettingSheet.vue')
  },
  {
    path: '/management',
    name: 'Management',
    component: () => import('../views/Management.vue')
  },
  {
    path: '/periodSeparatePage',
    name: 'PeriodSeparatePage',
    component: () => import('../views/PeriodSeparatePage.vue')
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
