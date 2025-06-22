import './assets/main.css'

import { createApp, type ComponentPublicInstance } from 'vue'
import App from './App.vue'
import router from './router'
import zitadelAuth from './lib/zitadel'
import { fatalError } from './lib/error'

async function main() {
  const authOk = await zitadelAuth.oidcAuth.startup()
  if (!authOk) {
    throw new Error('Auth startup not ok')
  }

  const app = createApp(App)
  app.use(router)
  app.config.errorHandler = errorHandler
  app.config.globalProperties.$zitadel = zitadelAuth
  app.mount('#app')
}

function errorHandler(
  err: unknown,
  instance: ComponentPublicInstance | null,
  info: string
) {
  const error = err instanceof Error ? err : new Error(String(err))
  fatalError(error)
}

main().catch(err => {
  console.error(err)
  const app = document.getElementById('app')
  const pre = document.createElement('pre')
  pre.textContent = String(err)
  app?.appendChild(pre)
})
