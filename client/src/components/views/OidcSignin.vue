<script setup lang="ts">
import { authCallback } from '@/lib/auth';
import { onMounted } from 'vue'
import { useRouter } from 'vue-router';

const router = useRouter()

onMounted(async () => {
  await authCallback(router)
})

</script>

<template>
  <div id="signin">
    <h1>Signing in</h1>
  </div>
</template>
