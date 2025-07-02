<script setup lang="ts">
import { auth } from '@/lib/auth';
import { onMounted } from 'vue'
import { useRouter } from 'vue-router';
import Status from '../Status.vue';

const router = useRouter()

onMounted(async () => {
  const done = await auth()
  if (done) {
    router.replace('/')
  }
})

</script>

<template>
  <Status title="Authorizing..." />
</template>
