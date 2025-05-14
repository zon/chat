<script setup lang="ts">

</script>

<template>
  <form id="new-message" action="/" method="post">
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