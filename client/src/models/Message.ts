import { post } from '@/lib/http'
import type { User } from './User'

export interface Message {
  ID: number
  User: User
  Content: string
  CreatedAt: string
  UpdatedAt: string
}

export async function postMessage(content: string) {
  const res = await post('messages', content)
  const data = await res.json()

  console.log('postMessage', data)

  return data as Message
}
