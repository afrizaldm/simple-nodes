// import { createRouter, createWebHistory } from 'vue-router'

import { createRouter, createWebHistory } from 'vue-router';

import Home from "../pages/Home.vue";


const routes: any[] = [
    {
        path: '/',
        name: 'index',
        redirect: { name: 'app' }
    },
    {
        path: '/app',
        name: 'app',
        component: Home
    }
]

const router = createRouter({
    history: createWebHistory('/app'),
    routes,
})

export default router