package game

import (
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
	h "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func ValidateKnightMove(fen, secondFEN *Fen) (bool, error) {
	move, err := h.LocateMove(fen, secondFEN)
	if err != nil {
		return false, err
	}
	pin, err := h.CheckPinedPeice(fen, *move)
	if err != nil {
		return false, nil
	}
	if !pin {
		return false, err
	}
	legal, err := ValidateKnightMove(fen, secondFEN)
	if err != nil {
		return false, err
	}
	if !legal {
		return false, nil
	}
	return true, nil
}
