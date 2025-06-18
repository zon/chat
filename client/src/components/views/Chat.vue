<script setup lang="ts">
import { onMounted } from 'vue'
import MessageView from '@/components/Message.vue'
import NewMessageForm from '@/components/NewMessageForm.vue'
import { sendMessage, updateMessages, messages } from '@/models/Message'
import { authUser } from '@/models/User'

onMounted(async () => {
  await updateMessages()
})

async function onNewMessage(content: string) {
  await sendMessage(content)
}

</script>

<template>
  <div id="chat">
    <div id="head">
      <div id="menu">
          <h1 id="title">Wurbs!</h1>
          <p>
            <RouterLink id="user" class="button" to="/auth">{{ authUser.name }}</RouterLink>
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
