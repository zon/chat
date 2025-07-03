import { AuthUser, type UserData } from '@/models/User'
import { ref } from 'vue'
import { get, put } from './http'
import { closeNats, connectNats } from './nats'
import { User, UserManager, WebStorageStateStore } from 'oidc-client-ts'
import type { Router } from 'vue-router'

const fullScope = import.meta.env.VITE_ZITADEL_FULL_SCOPE === 'true'
const zitadelProjectId = import.meta.env.VITE_ZITADEL_PROJECT_ID
const scope = [
  'openid profile email offline_access',
  ...(fullScope ? [
    `urn:zitadel:iam:org:project:id:${zitadelProjectId}:aud`,
    'urn:zitadel:iam:org:project:roles'
  ] : [])
].join(' ')

let oidcUser: User | null = null

export const authUser = ref(new AuthUser())

const authManager = new UserManager({
  response_type: 'code',
  scope,
  authority: import.meta.env.VITE_ZITADEL_ISSUER,
  client_id: import.meta.env.VITE_ZITADEL_CLIENT_ID,
  redirect_uri: `${window.location.origin}/oidc/signin`,
  userStore: new WebStorageStateStore({ store: localStorage })
})

export async function auth() {
  oidcUser = await authManager.getUser()
  if (oidcUser === null) {
    await authManager.signinRedirect()
    return false
  }
  await getAuthUser()
  const u = authUser.value
  console.debug('previous session', u.id, u.name)
  return true
}

export function isAuthed() {
  return oidcUser !== null && !authUser.value.isEmpty()
}

export async function authCallback(router: Router) {
  const user = await authManager.signinCallback()
  oidcUser = user ?? null
  await authManager.storeUser(oidcUser)
  await getAuthUser()
  const u = authUser.value
  console.debug('new session', u.id, u.name)
  router.replace('/')
}

export function getAccessToken() {
  return oidcUser?.access_token ?? null
}

export async function clearAuth() {
  await authManager.removeUser()
  oidcUser = null
  authUser.value = new AuthUser()
  await closeNats()
}

export async function renameAuthUser(name: string) {
  if (authUser.value.isEmpty()) {
    return authUser
  }
  const data = await put<UserData>(authUser.value.path(), { Name: name })
  authUser.value = new AuthUser(data)
  return authUser
}

async function getAuthUser() {
  const data = await get<UserData>('auth')
  authUser.value = new AuthUser(data)
  await connectNats()
}
