
// TODO consider dice roll
export const CheckRules = (store, squareIdx, circleIdx) => {

    if (store.Player.value.game.game_board.squares[squareIdx].captured_by != -1 && store.Player.value.game.game_board.squares[squareIdx].captured_by != store.Player.value.id) {
        return { allowed: false, reason: "You can't capture a circle in this square because the square is already captured by another player" };
    } else if (store.Player.value.game.game_board.squares[squareIdx].captured_by == store.Player.value.id) {
        return { allowed: false, reason: "You have already captured this circle, pick another circle" };
    } else {
        if (store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != -1 && store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != store.Player.value.id) {
            return { allowed: false, reason: "You can't select this circle because it is already selected by another player" };
        }
        return { allowed: true, reason: "You captured the circle" };
    }
}

