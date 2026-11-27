package game

import (
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
	h "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func ValidateQueenMove(fen, secondFEN *Fen) (bool, error) {
	move, err := h.LocateMove(fen, secondFEN)
	if err != nil {
		return false, err
	}
	notPined, err := h.CheckPinedPeice(fen, *move)
	if err != nil {
		return false, err
	}
	if !notPined {
		return false, nil
	}
	var legalMove bool

	if move.From.File == move.To.File || move.From.Rank == move.To.Rank {
		legalMove, err = h.ValidateVirHorzMoves(*move, *fen)
		if err != nil {
			return false, err
		}
	} else {
		legalMove, err = h.ValidateDiagnalMoves(*move, *fen)
		if err != nil {
			return false, err
		}
	}

	if !legalMove {
		return false, nil
	}
	return true, nil
}
