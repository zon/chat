import type { User } from '@/models/User'
import { ref } from 'vue'
import { get } from './http'

export const user = ref<User | null>(null)

export async function getAuth() {
  const u = await get<User>('auth')
  user.value = u
  return u
}
