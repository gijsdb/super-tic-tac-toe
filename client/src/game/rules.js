
export const CheckRules = (store, squareIdx, circleIdx, totalRoll) => {

    if (store.Player.value.game.game_board.squares[squareIdx].captured_by != -1 && store.Player.value.game.game_board.squares[squareIdx].captured_by != store.Player.value.id) {
        return { allowed: false, reason: "You can't capture a circle in this square because the square is already captured by another player" };
    } else if (store.Player.value.game.game_board.squares[squareIdx].captured_by == store.Player.value.id) {
        return { allowed: false, reason: "You have already captured this circle, pick another circle" };
    } else {
        if (store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != -1 && store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != store.Player.value.id) {
            return { allowed: false, reason: "You can't select this circle because it is already selected by another player" };
        }


        // If roll === index of a square its any circle in that square but the middle
        if (totalRoll >= 3 && totalRoll <= 11 && squareIdx >= 3 && squareIdx <= 11 && circleIdx != 4) {
            console.log("index of a square its any circle in that square but the middle")
            return { allowed: true, reason: "You captured the circle" };
        } else if (circleIdx === 4) {
            return { allowed: false, reason: `Can't select the middle` };
        }

        // 12 = any circle and square
        if (totalRoll === 12) {
            console.log("any circle and square")
            return { allowed: true, reason: "You captured the circle" };
        }

        // 7 = in any square middle
        if (totalRoll === 7 && circleIdx === 4) {
            console.log("in any square middle")
            return { allowed: true, reason: "You captured the circle" };
        }
        return { allowed: true, reason: `You captured the circle` };

    }
}

