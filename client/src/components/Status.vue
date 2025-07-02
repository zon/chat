<script setup lang="ts">

const props = defineProps<{
  title: string
}>()

</script>

<template>
  <div id="status">
    <h1>{{ props.title }}</h1>
  </div>
</template>

<style scoped>
  #status {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }
</style>
