import { createRouter, createWebHistory, type Router } from 'vue-router'
import { authManager } from './lib/auth'
import { fatalError } from './lib/error'

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
        }
    ]

    router = createRouter({
        history: createWebHistory(),
        routes
    })

    router.addRoute({
        path: '/oidc/signin',
        component: {
            created() {
                authManager
                    .signinCallback()
                    .then(async user => {
                        console.debug('OIDC signin', user)
                        await authManager.storeUser(user ?? null)
                        router.replace('/')
                    })
                    .catch(fatalError)
            }
        }
    })

    return router
}
