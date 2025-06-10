import { HOST } from './config'
import { getAccessToken } from './zitadel'

export async function get<T>(url: string) {
  const fullUrl = `${HOST}/${url}`
  const res = await fetch(fullUrl, {
    method: 'GET',
    headers: authHeaders()
  })
  isOk(res)
  const data = await res.json()
  return data as T
}

export async function post<T>(url: string, body: any) {
  const fullUrl = `${HOST}/${url}`
  const res = await fetch(fullUrl, {
    method: 'POST',
    headers: {
      ...authHeaders(),
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(body)
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
