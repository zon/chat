import { get, put } from "@/lib/http"

export interface User {
  ID: number
  Name: string
  Ready: boolean
  CreatedAt: string
  UpdatedAt: string
}

const path = 'users'

export function getUser(id: number) {
  return get<User>(`${path}/${id}`)
}

export function putUser(id: number, name: string) {
  return put<User>(`${path}/${id}`, { Name: name })
}
