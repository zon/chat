<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { getMessages, postMessage, type Message } from '@/models/Message'
import MessageView from '@/components/Message.vue'
import NewMessageForm from '@/components/NewMessageForm.vue'
import { authUser } from '@/lib/auth'

const authUserUrl = computed(() => `/users/${authUser.value?.ID}`)
const messages = ref<Message[]>([])

onMounted(async () => {
  messages.value = await getMessages()
})

async function onNewMessage(content: string) {
  const message = await postMessage(content)
  messages.value.unshift(message)
}

</script>

<template>
  <div id="chat">
    <div id="head">
      <div id="menu">
          <h1 id="title">Wurbs!</h1>
          <p>
            <RouterLink id="user" class="button" :to="authUserUrl">{{ authUser?.Name }}</RouterLink>
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
