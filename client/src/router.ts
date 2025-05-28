import { createWebHashHistory, createRouter } from 'vue-router'
import Chat from './components/Chat.vue'
import User from './components/User.vue'

const routes = [
    { path: '/', component: Chat },
    { path: '/users/:id', component: User }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
