import { authUser } from '@/lib/auth'
import { get } from '@/lib/http'
import type { Msg, Subscription } from '@nats-io/nats-core'
import { ref, type Ref } from 'vue'

export const usersSubject = 'users'

export interface UserData {
  ID: number
  Name: string
  Ready: boolean
  CreatedAt: string
  UpdatedAt: string
}

export class User {
  id: number
  name: string

  constructor(data?: UserData) {
    this.id = data?.ID || 0
    this.name = data?.Name || ''
  }

  isEmpty() {
    return this.id === 0
  }

  path() {
    return `${path}/${this.id}`
  }

}

export class AuthUser extends User {
  ready: boolean
  createdAt: Date
  updatedAt: Date

  constructor(data?: UserData) {
    super(data)
    this.ready = data?.Ready || false
    this.createdAt = new Date(data?.CreatedAt || 0)
    this.updatedAt = new Date(data?.UpdatedAt || 0)
  }

}

const path = 'users'

const users: { [id: string]: Ref<User> } = {}

export function getUser(id: number): Ref<User> {
  let user: Ref<User>
  if (!users[id]) {
    user = ref(new User())
    users[id] = user
    get<UserData>(`${path}/${id}`).then(data => {
      user.value = new User(data)
    })
  } else {
    user = users[id]
  }
  return user
}

export async function onUser(msg: Msg) {
  const data = msg.json<UserData>()
  const user = new User(data)
  setUser(user)
  if (user.id === authUser.value.id) {
    authUser.value = new AuthUser(data)
  }
}

function setUser(user: User) {
  if (!users[user.id]) {
    users[user.id] = ref(user)
  }
  users[user.id].value = user
}
