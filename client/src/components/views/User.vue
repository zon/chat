<script setup lang="ts">
import { authUser } from '@/lib/auth'
import { getUser, putUser } from '@/models/User'
import router from '@/router'
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const id = computed(() => Number(route.params.id))
const name = ref<string>('')
const ready = ref<boolean>(false)
const loaded = ref<boolean>(false)

onMounted(async () => {
  loaded.value = false
  const user = await getUser(id.value)
  name.value = user.Name
  ready.value = user.Ready
  loaded.value = true
})

async function onSubmit() {
  const user = await putUser(id.value, name.value)
  name.value = user.Name
  ready.value = user.Ready
  authUser.value = user
  router.push('/')
}

</script>

<template>
  <div id="page">
    <h1 id="title">Wurbs!</h1>
    <div v-if="loaded" id="content">
      <h2 v-if="ready">Edit User #{{ id }}</h2>
      <div v-else>
        <h2>Welcome</h2>
        <p>Set your name. <span class="note">Can be changed at any time</span></p>
      </div>
      <form @submit.prevent="onSubmit">
        <div class="field">
          <label for="name">Name</label>
          <input v-if="ready" id="name" name="name" type="text" v-model="name" />
          <input v-else id="name" name="name" type="text" />
        </div>
        <div class="actions">
          <button v-if="ready" class="primary" type="submit">Save</button>
          <RouterLink v-if="ready" class="button" to="/">Cancel</RouterLink>
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
