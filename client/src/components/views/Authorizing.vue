<script setup lang="ts">
import { auth, authUser } from '@/lib/auth';
import { onMounted } from 'vue'
import { useRouter } from 'vue-router';
import Status from '../Status.vue';

const router = useRouter()

onMounted(async () => {
  const done = await auth()
  if (done) {
    if (authUser.value.ready) {
      router.replace('/')
    } else {
      router.replace('/auth')
    }
  }
})

</script>

<template>
  <Status title="Authorizing..." />
</template>
