<script setup lang="ts">
import { authUser, renameAuthUser } from '@/lib/auth'
import { BadRequestError } from '@/lib/http'
import { AuthUser } from '@/models/User'
import { router } from '@/router'
import { onMounted, ref, watch, type Ref } from 'vue'

const user = ref(new AuthUser())
const error: Ref<BadRequestError | null> = ref(null)

onMounted(() => {
  user.value = authUser.value
})

watch(authUser, (newAuthUser) => {
  user.value = newAuthUser
})

async function onSubmit() {
  try {
    await renameAuthUser(user.value.name)
  } catch (err) {
    if (err instanceof BadRequestError) {
      error.value = err
      return
    }
    error.value = null
    throw err
  }
  router.push('/')
}

</script>

<template>
  <div id="page">
    <h1 id="title">Wurbs!</h1>
    <div v-if="!user.isEmpty()" id="content">
      <h2 v-if="user.ready">Edit User #{{ user.id }}</h2>
      <div v-else>
        <h2>Welcome</h2>
        <p>Set your name. <span class="note">Can be changed at any time</span></p>
      </div>
      <form @submit.prevent="onSubmit">
        <div class="field">
          <label for="name">Name</label>
          <input v-if="user.ready" id="name" name="name" type="text" v-model="user.name" />
          <input v-else id="name" name="name" type="text" />
        </div>
        <div v-if="error" id="error">
          <p>{{ error.message }}</p>
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
