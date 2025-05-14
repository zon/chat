import type { User } from './User'

export interface Message {
  id: number
  user: User
  content: string
  createdAt: Date
  updatedAt: Date
}
