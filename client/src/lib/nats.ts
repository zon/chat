import { wsconnect, type NatsConnection } from '@nats-io/nats-core'
import { get } from './http'

export let nats: NatsConnection

export async function connectNats() {
  const credentials = await get<string>('websocket')
  nats = await wsconnect({ servers: [credentials] })
}

export function listen<T>(subject: string, callback: (msg: T) => void) {
  return nats.subscribe(subject, {
    callback: (err, msg) => {
      if (err) {
        console.error(`Error receiving message on subject ${subject}:`, err)
        return
      }
      let data: T
      try {
        data = msg.json<T>()
      } catch (err) {
        console.error(`Error parsing message on subject ${subject}:`, err)
        return
      }
      try {
        callback(data)
      } catch (err) {
        console.error(`Error handling message on subject ${subject}:`, err)
      }
    }
  })
}

