<script setup lang="ts">
import { error } from '@/lib/error'

async function onSubmit() {
  location.reload()
}

</script>

<template>
  <div v-if="error !== null" id="error">
    <div id="error-page">
      <form id="error-content" @submit.prevent="onSubmit">
        <h2>{{ error.name }}</h2>
        <p>{{ error.message }}</p>
        <p>{{ error.stack }}</p>
        <div class="actions">
          <button type="submit">Reload</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
  #error {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: hsla(0, 0%, 60%, 60%);
  }
  #error-page {
    margin: auto;
    padding-top: 5em;
    max-width: 500px;
  }
  #error-content {
    padding: 1em 2ex;
    background-color: hsl(0, 95%, 40%);
  }
  #error-content > *:first-child {
    margin-top: 0;
  }
  #error-content > *:last-child {
    margin-bottom: 0;
  }
  #error-content button {
    color: hsl(0, 95%, 35%);
    font-weight: 600;
    background-color: white;
  }
</style>
