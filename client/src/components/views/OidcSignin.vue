<script setup lang="ts">
import { authCallback } from '@/lib/auth';
import { onMounted } from 'vue'
import { useRouter } from 'vue-router';
import Status from '../Status.vue';

const router = useRouter()

onMounted(async () => {
  await authCallback(router)
})

</script>

<template>
  <Status title="Signing in..." />
</template>
