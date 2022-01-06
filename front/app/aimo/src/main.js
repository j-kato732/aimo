import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'

// Auth0で使用するコンフィグの取得
import { domain, clientId } from '../auth_config.json';

// Auth0プラグインの取得
import { Auth0Plugin } from './auth';

Vue.use(Auth0Plugin, {
  domain,
  clientId,
  onRedirectCallBack: appState => {
    router.push(
      appState && appState.targetUrl
        ? appState.targetUrl
        : window.location.pathname
    );
  }
});

Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
