import Vue from 'vue'
import VueRouter from 'vue-router'
import { authGuard } from '../auth/auth0Guard'
import MyPage from '../views/MyPage.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/mypage',
    name: 'MyPage',
    component: MyPage,
    beforeEnter: authGuard
  },
  {
    path: '/',
    name: 'HelloWorld',
    component: () => import('../components/HelloWorld.vue')
  },
  {
    path: '/setting_user',
    name: 'Setting_user',
    component: () => import('../views/Setting_user.vue'),
    beforeEnter: authGuard
  },
  {
    path: '/aimSettingSheet',
    name: 'AimSettingSheet',
    component: () => import('../views/AimSettingSheet.vue'),
    beforeEnter: authGuard
  },
  {
    path: '/management',
    name: 'Management',
    component: () => import('../views/Management.vue'),
    beforeEnter: authGuard
  },
  {
    path: '/periodSeparatePage',
    name: 'PeriodSeparatePage',
    component: () => import('../views/PeriodSeparatePage.vue'),
    beforeEnter: authGuard
  },
  {
    path: '/registration',
    name: 'Registration',
    component: () => import('../views/Registration.vue'),
    beforeEnter: authGuard
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
