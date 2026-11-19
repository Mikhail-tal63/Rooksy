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
