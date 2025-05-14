<script setup lang="ts">
import { computed } from 'vue'
import type { Message } from '@/models/Message'
import { formatDate } from '@/lib/date'

const props = defineProps<{
  message: Message
}>()

const id = computed(() => `message-${props.message.id}`)
const time = computed(() => formatDate(props.message.createdAt))
</script>

<template>
  <div :id class="message">
    <p class="details">
      <span class="user">{{ message.user.name }}</span> <span class="time">{{ time }}</span>
    </p>
    <div class="content" v-html="message.content"></div>
  </div> 
</template>

<style scoped>
  .message .details {
    color: hsl(0, 0%, 50%);
    margin: 0;
  }
  .message .details .user {
    color: hsl(60 90% 75%);
  }
  .message .content > *:first-child {
    margin-top: 0;
  }
  .message .content > *:last-child {
    margin-bottom: 0;
  }
  .message .content pre {
    background-color: hsl(0 0% 10%);
    padding: 1em 2ex;
    overflow-x: scroll;
  }
  .message .content code {
    background-color: hsl(0 0% 10%);
  }
</style>