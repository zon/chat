import { AuthUser, type UserData } from '@/models/User'
import { ref } from 'vue'
import { get, put } from './http'
import { connectNats } from './nats'

export const authUser = ref(new AuthUser())

export async function updateAuth() {
  if (authUser.value.isEmpty()) {
    const data = await get<UserData>('auth')
    authUser.value = new AuthUser(data)
  }
  await connectNats()
  return authUser
}

export function clearAuth() {
  authUser.value = new AuthUser()
}

export async function renameAuthUser(name: string) {
  if (authUser.value.isEmpty()) {
    return authUser
  }
  const data = await put<UserData>(authUser.value.path(), { Name: name })
  authUser.value = new AuthUser(data)
  return authUser
}