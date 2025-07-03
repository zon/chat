import { clearAuth, getAccessToken } from './auth'
import { HOST } from './config'

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

export async function get<T>(url: string, params?: { [key: string]: string }) {
  let searchUrl = url
  if (params !== undefined) {
    const search = new URLSearchParams(params)
    searchUrl = `${searchUrl}?${search}`
  }
  return request<T>(searchUrl, { method: 'GET' })
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
  } else if (res.status === 401) {
    await clearAuth()
    location.reload()
  }
  if (!res.ok) {
    throw new ResponseError(res, url, init)
  }
  return res
}

function authHeaders(): { Authorization?: string } {
  const token = getAccessToken()
  if (token === null) {
    return {}
  }
  const bearer = `Bearer ${token}`
  return { Authorization: bearer }
}
