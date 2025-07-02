<script setup lang="ts">
import { auth } from '@/lib/auth';
import { onMounted } from 'vue'
import { useRouter } from 'vue-router';

const router = useRouter()

onMounted(async () => {
  const done = await auth()
  if (done) {
    router.replace('/')
  }
})

</script>

<template>
  <div id="authorizing">
    <h1>Authorizing</h1>
  </div>
</template>
