// import { createRouter, createWebHistory } from 'vue-router'

import { createRouter, createWebHistory } from 'vue-router';

import Home from "../pages/Home.vue";


const routes: any[] = [
    {
        path: '/',
        name: 'index',
        component: Home
        // redirect: { name: 'home' }
    },
    {
        path: '/home',
        name: 'home',
        component: Home
    }
]

const router = createRouter({
    history: createWebHistory('/app'),
    routes,
})

export default router