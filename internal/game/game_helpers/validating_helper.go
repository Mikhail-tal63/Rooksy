package gamehelpers

import . "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"

func CalcolateMoves(firstFEN, secondFEN *Fen) (*Move, error) {
	var from, to Position
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if firstFEN.Board[i][j] != secondFEN.Board[i][j] {
				if secondFEN.Board[i][j].Type == Empty {
					from = Position{
						File: i,
						Rank: j,
					}
				} else {
					to = Position{
						File: i,
						Rank: j,
					}
				}
			}
		}
	}

	move := &Move{
		From: from,
		To:   to,
	}
	return move, nil
}

func ValidateVirHorzPeices(move Move, fen Fen) (bool, error) {
	var from, to Position = move.From, move.To
	board := fen.Board
	if from == to {

		return false, nil
	}

	if from.File == to.File {

		if from.Rank > to.Rank {
			for i := to.Rank + 1; i < from.Rank; i++ {
				if board[from.File][i].Type != Empty {

					return false, nil
				}
			}
		} else {
			for i := from.Rank + 1; i < to.Rank; i++ {
				if board[from.File][i].Type != Empty {

					return false, nil
				}
			}
		}

	} else if from.Rank == to.Rank {

		if from.File > to.File {
			for i := to.File + 1; i < from.File; i++ {
				if board[from.File][i].Type != Empty {

					return false, nil
				}
			}
		} else {
			for i := from.File + 1; i < to.File; i++ {
				if board[i][from.Rank].Type != Empty {

					return false, nil
				}
			}
		}

	} else {

		return false, nil
	}
	return true, nil

}
