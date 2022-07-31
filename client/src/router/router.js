import Home from "../views/Home.vue";
import Board from "../views/Board.vue";
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    { path: "/", component: Home },
    { path: "/game/:id", component: Board }
];

export const Router = createRouter({
    history: createWebHistory(),
    routes
})

Router.beforeEach((to, from) => {
    console.log("BEFORE EACH ROUTE")
})
