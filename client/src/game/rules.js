
export const CheckRules = (store, squareIdx, circleIdx, totalRoll) => {

    if (store.Player.value.game.game_board.squares[squareIdx].captured_by != -1 && store.Player.value.game.game_board.squares[squareIdx].captured_by != store.Player.value.id) {
        return { allowed: false, reason: "You can't capture a circle in this square because the square is already captured by another player" };
    } else if (store.Player.value.game.game_board.squares[squareIdx].captured_by == store.Player.value.id) {
        return { allowed: false, reason: "You have already captured this circle, pick another circle" };
    } else {
        if (store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != -1 && store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != store.Player.value.id && totalRoll != 2) {
            return { allowed: false, reason: "You can't select this circle because it is already selected by another player" };
        }

        // A throw of twelve gives the player the freedom to peg any hole. A throw of seven lets a player peg any box's center hole. 

        // A throw of two lets one remove an opponent's peg. A player wins a box by getting three in a row or any five holes.
        if (totalRoll === 2 && store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != -1 && store.Player.value.game.game_board.squares[squareIdx].circles[circleIdx].selected_by != store.Player.value.id && store.Player.value.game.game_board.squares[squareIdx].captured_by == -1) {
            console.log("remove others peg")
            return { allowed: true, reason: "You've removed your opponents circle!" }
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
        } else if (totalRoll === 7 && circleIdx !== 4) {
            return { allowed: false, reason: "You can't select this circle because it is not in the middle" };
        }

        // If player selects a circle that does not match the total roll value then they can't select it.
        if (totalRoll !== circleIdx + 3 && totalRoll != 2) {
            return { allowed: false, reason: `You can only select a circle with the number ${totalRoll}` };
        } else if (totalRoll !== circleIdx + 3 && totalRoll == 2) {
            return { allowed: false, reason: `You can only remove a circle from the opposition with the number ${totalRoll}` };
        } else {
            return { allowed: true, reason: "You captured the circle" };
        }
    }
}

