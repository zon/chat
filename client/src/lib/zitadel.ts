import { createZITADELAuth } from '@zitadel/vue'
import { SignInType, type OidcAuth } from 'vue-oidc-client/vue3'
import { fatalError } from './error'
import { clearAuth, updateAuth } from './auth'
import { WebStorageStateStore } from 'oidc-client'

export interface ZitadelAuth {
  oidcAuth: OidcAuth
  hasRole: (role: string) => any
}

export let zitadelAuth: ZitadelAuth

export async function initZitadel() {

  zitadelAuth = createZITADELAuth({
    project_resource_id: import.meta.env.VITE_ZITADEL_PROJECT_RESOURCE_ID,
    client_id: import.meta.env.VITE_ZITADEL_CLIENT_ID,
    issuer: import.meta.env.VITE_ZITADEL_ISSUER,
  }, undefined, SignInType.Window, undefined, {
    automaticSilentRenew: true,
    stateStore: new WebStorageStateStore({ store: localStorage }),
  })

  zitadelAuth.oidcAuth.events.addSilentRenewError(fatalError)

  zitadelAuth.oidcAuth.events.addUserLoaded(_ => {
    updateAuth().catch(fatalError)
  })

  zitadelAuth.oidcAuth.events.addUserUnloaded(clearAuth)
}

export async function startZitadel() {
  const ok = await zitadelAuth.oidcAuth.startup()
  if (!ok) {
    throw new Error('Auth startup not ok')
  }
  if (zitadelAuth.oidcAuth.isAuthenticated) {
    await updateAuth()
  }
}

export function getAuthProfile() {
  return zitadelAuth.oidcAuth.userProfile
}

export function getAccessToken() {
  return zitadelAuth.oidcAuth.accessToken
}
