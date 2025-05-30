/// <reference types="vite/client" />

import zitadelAuth from './lib/zitadel'

declare module 'vue' {
  interface ComponentCustomProperties {
    $zitadel: typeof zitadelAuth
  }
}
