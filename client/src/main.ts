import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import { initRouter } from './router'
import { initZitadel, startZitadel, zitadelAuth } from './lib/zitadel'
import { fatalError } from './lib/error'
import { connectNats } from './lib/nats'

async function main() {
  await initZitadel()

  const router = initRouter()

  const app = createApp(App)
  app.use(router)
  app.config.errorHandler = errorHandler
  app.config.globalProperties.$zitadel = zitadelAuth
  app.mount('#app')

  try {
    await startZitadel()
  } catch (err) {
    errorHandler(err)
  }
}

function errorHandler(err: unknown) {
  if (err instanceof Error) {
    fatalError(err)
  } else {
    fatalError(new Error(String(err)))
  }
}

main().catch(err => {
  console.error(err)
  const app = document.getElementById('app')
  const pre = document.createElement('pre')
  pre.textContent = String(err)
  app?.appendChild(pre)
})
