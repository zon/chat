<script setup lang="ts">
import { onMounted } from 'vue'
import { getAuth, subscribeUsers } from './models/User'
import { connectNats } from './lib/nats'
import { subscribeMessages } from './models/Message'
import FatalError from './components/FatalError.vue'

onMounted(async () => {
  await getAuth()
  await connectNats()
  subscribeUsers()
  subscribeMessages()
})
</script>

<template>
  <RouterView />
  <FatalError />
</template>

<style scoped>

</style>
