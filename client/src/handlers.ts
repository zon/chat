import type { Nats } from './lib/nats'
import { messagesSubject, onMessage, onMessageReconnect } from './models/Message'
import { usersSubject, onUser, onUserReconnect } from './models/User'

export function addHandlers(nats: Nats) {
  nats.subscribe(messagesSubject, onMessage, onMessageReconnect)
  nats.subscribe(usersSubject, onUser, onUserReconnect)
}
