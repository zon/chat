import { HOST } from './config'
import { getAccessToken } from './zitadel'

export function get(url: string) {
  const fullUrl = `${HOST}/${url}`
  const token = getAccessToken()
  const bearer = `Bearer ${token}`
  return fetch(fullUrl, {
    method: 'GET',
    headers: {
      Authorization: bearer
    }
  })
}

export function post(url: string, body: any) {
  const fullUrl = `${HOST}/${url}`
  const token = getAccessToken()
  const bearer = `Bearer ${token}`
  return fetch(fullUrl, {
    method: 'POST',
    headers: {
      Authorization: bearer,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(body)
  })
}
