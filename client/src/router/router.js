import Home from "../views/Home.vue";
import Board from "../views/Board.vue";
import { createRouter, createWebHistory } from 'vue-router';
import { useGameStore } from "../stores/game.js";
import { storeToRefs } from "pinia";


const routes = [
    { name: 'Home', path: "/", component: Home },
    {
        name: 'Game', path: "/game/:id", component: Board,
        beforeEnter: async (to, from, next) => {
            const store = useGameStore();
            let playerStore = storeToRefs(store);
            if (!playerStore.Player.value.inGame && playerStore.Player.value.game.ID != to.params.id) {
                next({ name: 'Home' })
                return
            } else {
                next()
                return
            }
        }
    }
];

export const Router = createRouter({
    history: createWebHistory(),
    routes
})

