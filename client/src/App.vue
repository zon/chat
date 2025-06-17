<script setup lang="ts">
import { onMounted } from 'vue'
import { getAuth, subscribeUsers } from './models/User'
import { connectNats } from './lib/nats'
import { subscribeMessages } from './models/Message'

onMounted(async () => {
  await getAuth()
  await connectNats()
  subscribeUsers()
  subscribeMessages()
})
</script>

<template>
  <RouterView />
</template>

<style scoped>

</style>
