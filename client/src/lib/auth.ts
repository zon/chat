import type { User } from '@/models/User'
import { ref } from 'vue'
import { get } from './http'

export const authUser = ref<User | null>(null)

export async function getAuth() {
  const u = await get<User>('auth')
  authUser.value = u
  return u
}
