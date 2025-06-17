import { get, post } from '@/lib/http'
import { getUser } from './User'
import { reactive } from 'vue'
import type { Subscription } from '@nats-io/nats-core'
import { listen } from '@/lib/nats'

interface MessageData {
  ID: number
  UserID: number
  Content: string
  CreatedAt: string
  UpdatedAt: string
}

export class Message {
  id: number
  userId: number
  content: string
  createdAt: Date
  updatedAt: Date

  constructor(data: MessageData) {
    this.id = data.ID
    this.userId = data.UserID
    this.content = data.Content
    this.createdAt = new Date(data.CreatedAt)
    this.updatedAt = new Date(data.UpdatedAt)
  }

  getUser() {
    return getUser(this.userId)
  }

}

const path = 'messages'

let messagesSubscription: Subscription

export const messages = reactive<Message[]>([])

export function subscribeMessages() {
  messagesSubscription = listen<MessageData>('messages', data => {
    addMessage(new Message(data))
  })
}

export async function updateMessages() {
  const list = await get<MessageData[]>(path)
  for (const data of list) {
    addMessage(new Message(data))
  }
}

export async function sendMessage(content: string) {
  const data = await post<MessageData>(path, content)
  return new Message(data)
}

function addMessage(message: Message) {
  for (let i = 0; i < messages.length; i++) {
    const other = messages[i]
    if (other.id === message.id) {
      messages[i] = message
      return
    }
    if (other.createdAt < message.createdAt) {
      messages.splice(i, 0, message)
      return
    }
  }
  messages.push(message)
}
