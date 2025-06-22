import { HOST } from './config'
import { getAccessToken } from './zitadel'

export class ResponseError extends Error {
  res: Response
  method: string
  url: URL
  init?: RequestInit

  constructor(
    res: Response,
    url: URL,
    init?: RequestInit
  ) {
    const method = init?.method || 'GET'
    super(`${res.status} response ${method} ${url}`)
    this.res = res
    this.method = method
    this.url = url
    this.init = init
  }

}

export class BadRequestError extends ResponseError {
  code: string

  constructor(
    bad: BadRequest,
    res: Response,
    url: URL,
    init?: RequestInit
  ) {
    super(res, url, init)
    this.code = bad.Code
    this.message = bad.Message
  }

}

export interface BadRequest {
  Code: string
  Message: string
}

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
  const fullUrl = new URL(`${HOST}/${url}`)
  const res = await fetchOkay(fullUrl, {
    ...init,
    headers: {
      ...authHeaders(),
      ...init.headers
    }
  })
  const data = await res.json()
  return data as T
}

async function fetchOkay(url: URL, init?: RequestInit) {
  const res = await fetch(url, init)
  if (res.status === 400) {
    const bad = await res.json()
    throw new BadRequestError(bad as BadRequest, res, url, init)
  }
  if (!res.ok) {
    throw new ResponseError(res, url, init)
  }
  return res
}

function authHeaders() {
  const token = getAccessToken()
  const bearer = `Bearer ${token}`
  return { Authorization: bearer }
}
