import { HOST } from './config'
import { getAccessToken } from './zitadel'

export async function get<T>(url: string) {
  return request<T>(url, { method: 'GET' })
}

export async function post<T>(url: string, body: any) {
  return request<T>(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
}

export async function put<T>(url: string, body: any) {
  return request<T>(url, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(body)
  })
}

async function request<T>(url: string, init: RequestInit) {
  const fullUrl = `${HOST}/${url}`
  const res = await fetch(fullUrl, {
    ...init,
    headers: {
      ...authHeaders(),
      ...init.headers
    }
  })
  isOk(res)
  const data = await res.json()
  return data as T
}

function authHeaders() {
  const token = getAccessToken()
  const bearer = `Bearer ${token}`
  return { Authorization: bearer }
}

function isOk(res: Response) {
  if (res.ok) {
    return true
  }
  throw new Error(`HTTP error! status: ${res.status}`)
}
