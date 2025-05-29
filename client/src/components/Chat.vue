<script setup lang="ts">
import { computed, ref } from 'vue'
import type { User } from '@/models/User'
import type { Message } from '@/models/Message'
import MessageView from '@/components/Message.vue'
import NewMessageForm from './NewMessageForm.vue'

const user: User = {
  id: 1,
  name: 'Zon',
  ready: true
}
const createdAt = new Date()
const updatedAt = createdAt
const messages = ref<Message[]>([
  {
    id: 2,
    content: '<p>World</p>',
    user,
    createdAt,
    updatedAt
  },
  {
    id: 1,
    content: '<p>Hello</p>',
    user,
    createdAt,
    updatedAt
  },
])

function onNewMessage(content: string) {
  const createdAt = new Date()
  const updatedAt = createdAt
  const message: Message = {
    id: messages.value.length + 1,
    content,
    user,
    createdAt,
    updatedAt
  }
  messages.value.unshift(message)
}

const userUrl = computed(() => `/users/${user.id}`)

</script>

<template>
  <div id="chat">
    <div id="head">
      <div id="menu">
          <h1 id="title">Wurbs!</h1>
          <p>
            <RouterLink id="user" class="button" :to="userUrl">{{ user.name }}</RouterLink>
          </p>
      </div>
    </div>
    <div id="messages">
      <MessageView v-for="message in messages" :message />
    </div>
    <div id="foot">
      <NewMessageForm @submit="onNewMessage" />
    </div>
  </div>
</template>

<style scoped>
  #chat {
    display: flex;
    flex-direction: column;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
  }
  #menu {
    display: flex;
    align-items: center;
  }
  #menu p {
    margin: 0;
  }
  #menu p a.button {
    border-width: 0 0 0 1px;
    border-style: solid;
    border-color: hsl(0 0% 30%);
    padding: 1em 2ex;
  }
  #menu h1#title {
    flex: 1;
  }

  #messages {
    padding: 1em 2ex;
    background-color: hsl(0 0% 15%);
  }
  #messages {
    display: flex;
    flex: 1;
    flex-direction: column-reverse;
    gap: 1em;
    overflow-y: scroll;
  }
  #foot {
    background-color: hsl(0 0% 20%);
  }
</style>
