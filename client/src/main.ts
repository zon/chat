import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import zitadelAuth from './lib/zitadel'

zitadelAuth.oidcAuth.startup().then(ok => {
  if (!ok) {
    console.error('Startup not ok')
    return
  }
  const app = createApp(App)
  app.use(router)
  app.config.globalProperties.$zitadel = zitadelAuth
  app.mount('#app')
})
