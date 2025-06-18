<script setup lang="ts">
import { authUser, renameAuthUser } from '@/models/User'
import router from '@/router'

async function onSubmit() {
  await renameAuthUser(authUser.value.name)
  router.push('/')
}

</script>

<template>
  <div id="page">
    <h1 id="title">Wurbs!</h1>
    <div id="content">
      <h2 v-if="authUser.ready">Edit User #{{ authUser.id }}</h2>
      <div v-else>
        <h2>Welcome</h2>
        <p>Set your name. <span class="note">Can be changed at any time</span></p>
      </div>
      <form @submit.prevent="onSubmit">
        <div class="field">
          <label for="name">Name</label>
          <input v-if="authUser.ready" id="name" name="name" type="text" v-model="authUser.name" />
          <input v-else id="name" name="name" type="text" />
        </div>
        <div class="actions">
          <button v-if="authUser.ready" class="primary" type="submit">Save</button>
          <RouterLink v-if="authUser.ready" class="button" to="/">Cancel</RouterLink>
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
