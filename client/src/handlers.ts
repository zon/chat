import type { Nats } from './lib/nats'
import { messagesSubject, onMessage } from './models/Message'
import { usersSubject, onUser } from './models/User'

export function addHandlers(nats: Nats) {
  nats.subscribe(messagesSubject, onMessage)
  nats.subscribe(usersSubject, onUser)
}
