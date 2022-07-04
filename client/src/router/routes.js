import Home from "../views/Home.vue";
import Board from "../views/Board.vue";

export const routes = [
    { path: "/", component: Home },
    { path: "/game", component: Board }
];
