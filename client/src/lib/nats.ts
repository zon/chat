import { wsconnect, type Msg, type NatsConnection, type Status, type Subscription } from '@nats-io/nats-core'
import { get } from './http'
import { addHandlers } from '@/handlers'
import { ref, type Ref } from 'vue'

const disconnectMargin = 2000

export let nats: Nats

interface Credentials {
  Host: string
  User: string
  Password: string
}

export async function connectNats() {
  if (nats !== undefined) {
    await nats.drain()
  }
  nats = await Nats.connect()
  addHandlers(nats)
  return nats
}

export async function closeNats() {
  if (nats !== undefined && !nats.isClosed()) {
    await nats.drain()
  }
}

export class Nats {
  status: Ref<Status>
  reconnecting: Ref<boolean>
  disconnected: Date

  private conn: NatsConnection
  private handlers: Handler[]
  private statusLoop: Promise<void> | null

  constructor(conn: NatsConnection) {
    this.status = ref({type: 'close'})
    this.reconnecting = ref(false)
    this.disconnected = new Date()
    this.conn = conn
    this.handlers = []
    this.statusLoop = this.startStatus()
  }

  static async connect() {
    const credentials = await get<Credentials>('websocket')
    const conn = await wsconnect({
      servers: [credentials.Host],
      user: credentials.User,
      pass: credentials.Password
    })
    return new this(conn)
  }

  isClosed() {
    return this.conn.isClosed()
  }

  subscribe(subject: string, onMsg: (msg: Msg) => Promise<void>, onReconnect: (disconnected: Date) => Promise<void>) {
    var sub = this.conn.subscribe(subject)
    this.handlers.push(new Handler(sub, onMsg, onReconnect))
  }

  private async startStatus() {
    for await (const status of this.conn.status()) {
      console.info('nats status', status.type)
      this.status.value = status
      const connected = (
        status.type !== 'reconnect' &&
        status.type !== 'ping'
      )
      if (connected === this.reconnecting.value) {
        continue
      }
      this.reconnecting.value = connected
      const since = new Date(this.disconnected.getTime() - disconnectMargin)
      if (connected) {
        for (const handler of this.handlers) {
          await handler.onReconnect(since)
        }
      } else {
        this.disconnected = new Date()
      }
    }
  }

  async drain() {
    await this.conn.drain()
    for (const handler of this.handlers) {
      await handler.loop
    }
    await this.statusLoop
  }

}

class Handler {
  sub: Subscription
  onMsg: (msg: Msg) => Promise<void>
  onReconnect: (disconnected: Date) => Promise<void>
  loop: Promise<void>

  constructor(sub: Subscription, onMsg: (msg: Msg) => Promise<void>, onReconnect: (disconnected: Date) => Promise<void>) {
    this.sub = sub
    this.onMsg = onMsg
    this.onReconnect = onReconnect
    this.loop = this.msgLoop()
  }

  private async msgLoop() {
    for await (const msg of this.sub) {
      console.debug('msg', this.sub.getSubject())
      try {
        await this.onMsg(msg)
      } catch (err) {
        console.error(`${this.sub} msg error:`, err)
      }
    }
  }

}
