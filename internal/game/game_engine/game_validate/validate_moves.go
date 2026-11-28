package game

import (
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_errors"
	h "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func ValidateBishopMove(fen, secondFEN *Fen) (bool, error) {
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
	legalMove, err := h.ValidateDiagnalMoves(*move, *fen)
	if err != nil {
		return false, err
	}
	if !legalMove {
		return false, nil
	}
	return true, nil
}
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
	legal, err := h.ValidateLShapedMoves(*move, *fen)
	if err != nil {
		return false, err
	}
	if !legal {
		return false, nil
	}
	return true, nil
}
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
func ValidatePawnMove(fen, secondFEN *Fen) (bool, error) {
	move, err := h.LocateMove(fen, secondFEN)
	if err != nil {
		return false, err
	}
	pin, err := h.CheckPinedPeice(fen, *move)
	if err != nil {
		return false, nil
	}
	if !pin {
		return false, nil
	}

	legal, err := h.ValidateLegalPawnMoves(*move, *fen)
	if err != nil {
		return false, err
	}
	if !legal {
		return false, nil
	}
	return true, nil
} /*
func ValidateKingMove(fen, secondFEN *Fen) (bool, error) {
	move, err := h.LocateMove(fen, secondFEN)
	if err != nil {
		return false, err
	}
	return true, nil
}*/
