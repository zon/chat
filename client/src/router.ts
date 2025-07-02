import { createRouter, createWebHistory, type Router } from 'vue-router'
import { authCallback, isAuthed } from './lib/auth'
import { fatalError } from './lib/error'

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
        }
    ]

    router = createRouter({
        history: createWebHistory(),
        routes
    })

    router.beforeEach(async (to, from) => {
        if (!to.meta.noAuth && !isAuthed()) {
            return { name: authorizing }
        }
    })

    router.addRoute({
        path: '/oidc/signin',
        component: {
            created() {
                authCallback(router).catch(fatalError)
            }
        },
        meta: { noAuth: true }
    })

    return router
}
