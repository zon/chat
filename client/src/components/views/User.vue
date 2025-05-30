<script setup lang="ts">
import type { User } from '@/models/User'

const user: User = {
  id: 1,
  name: 'Zon',
  ready: true
}

</script>

<template>
  <div id="page">
    <h1 id="title">Wurbs!</h1>
    <div id="content">
      <h2 v-if="user.ready">Edit User #{{ user.id }}</h2>
      <div v-else>
        <h2>Welcome</h2>
        <p>Set your name. <span class="note">Can be changed at any time</span></p>
      </div>
      <form>
        <div class="field">
          <label for="name">Name</label>
          <input id="name" name="name" type="text" :value="user.name" />
        </div>
        <div class="actions">
          <button v-if="user.ready" class="primary" type="submit">Save</button>
          <RouterLink v-if="user.ready" class="button" to="/">Cancel</RouterLink>
          <button v-else class="primary" type="submit">Set</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
  #page {
    margin: auto;
    max-width: 500px;
  }
  #content {
    border: 1px solid hsl(0 0% 30%);
    padding: 1em 2ex;
  }
  #content > *:first-child {
    margin-top: 0;
  }
  #content > *:last-child {
    margin-bottom: 0;
  }
</style>
