import { get, post } from '@/lib/http'
import type { User } from './User'

export interface Message {
  ID: number
  User: User
  Content: string
  CreatedAt: string
  UpdatedAt: string
}

const path = 'messages'

export function getMessages() {
  return get<Message[]>(path)
}

export function postMessage(content: string) {
  return post<Message>(path, content)
}
