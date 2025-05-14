<script setup lang="ts">
const emit = defineEmits<{
  submit: [content: string]
}>()

function onSubmit() {
  const contentElement = document.getElementById('new-message-content') as HTMLParagraphElement
  const content = contentElement.innerText
  if (content) {
    emit('submit', content)
    contentElement.innerText = ''
  }
}
</script>

<template>
  <form id="new-message" @submit.prevent="onSubmit">
    <p id="new-message-content" contenteditable="plaintext-only"></p>
    <button class="primary" type="submit">📨</button>
  </form>
</template>

<style scoped>
  #new-message {
    display: flex;
  }
  #new-message-content {
    flex: 1;
    padding: 1em 2ex;
    margin: 0;
    content: "Message";
  }
  #new-message-content:focus {
    outline: none;
    content: normal;
  }
</style>