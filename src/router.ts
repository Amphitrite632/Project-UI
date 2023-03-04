import { createRouter, createWebHistory } from "vue-router"
import Home from "./pages/Home.vue"
import About from "./pages/About.vue"

const routes = [
    {
        path: "/",
        name: "home",
        component: Home,
    },
    {
        path: "/about",
        name: "about",
        component: About,
    },
    {
        path: "/edit/illust",
        name: "home",
        component: Home,
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router
