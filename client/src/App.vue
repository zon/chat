<script setup lang="ts">
import { onMounted } from 'vue'
import { getAuth } from './lib/auth'

onMounted(async () => {
  await getAuth()
})
</script>

<template>
  <RouterView />
</template>

<style scoped>

</style>
