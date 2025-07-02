import { createRouter, createWebHistory, type Router } from 'vue-router'
import { isAuthed } from './lib/auth'

const authorizing = 'authorizing'

export let router: Router

export function initRouter() {
    const routes = [
        {
            path: '/',
            component: () => import('./components/views/Chat.vue')
        },
        {
            path: '/auth',
            component: () => import('./components/views/User.vue')
        },
        {
            name: authorizing,
            path: '/authorizing',
            component: () => import('./components/views/Authorizing.vue'),
            meta: { noAuth: true }
        },
        {
            path: '/oidc/signin',
            component: () => import('./components/views/OidcSignin.vue'),
            meta: { noAuth: true }
        }
    ]

    router = createRouter({
        history: createWebHistory(),
        routes
    })

    router.beforeEach(to => {
        if (!to.meta.noAuth && !isAuthed()) {
            return { name: authorizing }
        }
    })

    return router
}
