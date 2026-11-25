package game

import (
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_errors"
	h "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func ValidateRookMove(fen, secondFEN *Fen) (bool, error) {
	move, err := h.LocateMove(fen, secondFEN)
	if err != nil {
		return false, err
	}
	notPined, err := h.CheckPinedPeice(fen, *move)
	if err != nil {
		return false, err
	}
	if !notPined {
		return false, ErrIllegalMove
	}
	legalMove, err := h.ValidateVirHorzMoves(*move, *fen)
	if err != nil {
		return false, err
	}
	if !legalMove {
		return false, ErrIllegalMove
	}
	return true, nil
}

func ValidRookMoves(fen *Fen) (*[]RookMoves, error) {
	return nil, nil
}
