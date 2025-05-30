import { createRouter, createWebHistory } from 'vue-router'
import zitadelAuth from './lib/zitadel'

const routes = [
    {
        path: '/',
        meta: {
            authName: zitadelAuth.oidcAuth.authName
        },
        component: () => import('./components/views/Chat.vue')
    },
    {
        path: '/users/:id',
        meta: {
            authName: zitadelAuth.oidcAuth.authName
        },
        component: () => import('./components/views/User.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

zitadelAuth.oidcAuth.useRouter(router)

export default router
