import { createRouter, createWebHistory, type Router } from 'vue-router'
import { zitadelAuth } from './lib/zitadel'

export let router: Router

export function initRouter() {
    const routes = [
        {
            path: '/',
            meta: {
                authName: zitadelAuth.oidcAuth.authName
            },
            component: () => import('./components/views/Chat.vue')
        },
        {
            path: '/auth',
            meta: {
                authName: zitadelAuth.oidcAuth.authName
            },
            component: () => import('./components/views/User.vue')
        }
    ]

    router = createRouter({
        history: createWebHistory(),
        routes
    })

    zitadelAuth.oidcAuth.useRouter(router)

    return router
}
