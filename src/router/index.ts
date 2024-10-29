// import { createRouter, createWebHistory } from 'vue-router'

import { createRouter, createWebHistory } from 'vue-router';

import Home from "../pages/Home.vue";
import Add from "../pages/Add.vue";
import Edit from "../pages/Edit.vue";

const routes: any[] = [
    {
        path: '/',
        name: 'index',
        component: Home
        // redirect: { name: 'home' }
    },
    {
        path: '/add/:id',
        name: 'add',
        component: Add
    },
    {
        path: '/edit/:id',
        name: 'edit',
        component: Edit
    }
]

const router = createRouter({
    history: createWebHistory('/app'),
    routes,
})

export default router