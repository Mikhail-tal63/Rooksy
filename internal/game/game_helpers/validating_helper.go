package gamehelpers

import (
	"fmt"

	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
)

// returns move the have been made
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

// checks if piece can go to certin squire
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
func CheckPinedPeice(fen *Fen, move Move) (bool, error) {
	from := move.From
	to := move.To
	board := fen.Board
	color := fen.Board[from.File][from.Rank].Color
	king, err := GetKingPosition(fen, color)
	if err != nil {
		return false, fmt.Errorf("could not find king : validate_helper")
	}
	if from.File == king.File {
		r := king.File
		l := king.File
		pass := false
		if from.File > king.File {
			for i := king.File + 1; i < from.File; i++ {
				if board[i][king.Rank].Type != Empty {
					pass = true
					break
				}
			}
			if !pass {
				for i := from.File + 1; i < 8; i++ {
					if board[i][king.Rank].Type != Empty {
						if board[i][king.Rank].Type == Queen && board[i][king.Rank].Color != color {
							return false, nil
						}
						if board[i][king.Rank].Type == Rook && board[i][king.Rank].Color != color {
							return false, nil
						}
						break
					}
				}
			}
		}
	}

}

func CheckPinedPeiceHorVir(fen *Fen, move Move) (bool, error) {
	from := move.From

	board := fen.Board
	color := fen.Board[from.File][from.Rank].Color
	king, err := GetKingPosition(fen, color)
	if err != nil {
		return false, fmt.Errorf("could not find king : validate_helper")
	}
	if from.File == king.File {

		pass := false
		if from.File > king.File {
			for i := king.File + 1; i < from.File; i++ {
				if board[i][king.Rank].Type != Empty {
					pass = true
					break
				}
			}
			if !pass {
				for i := from.File + 1; i < 8; i++ {
					if board[i][king.Rank].Type != Empty {
						if board[i][king.Rank].Type == Queen && board[i][king.Rank].Color != color {
							return false, nil
						}
						if board[i][king.Rank].Type == Rook && board[i][king.Rank].Color != color {
							return false, nil
						}
						break
					}
				}
			}
		} else {
			pass = false
			for i := king.File - 1; i > from.File; i-- {
				if board[i][king.Rank].Type != Empty {
					pass = true
					break
				}
			}
			if !pass {
				for i := from.File - 1; i >= 0; i-- {
					if board[i][king.Rank].Type != Empty {
						if board[i][king.Rank].Type == Queen && board[i][king.Rank].Color != color {
							return false, nil
						}
						if board[i][king.Rank].Type == Rook && board[i][king.Rank].Color != color {
							return false, nil
						}
						break
					}
				}
			}
		}
	}

	return true, nil
}

func P_CheckRank(board [8][8]Piece, from Position, color Colors, king Position, to Position) (bool, error) {
	if from.File == king.File {
		if from.Rank > king.Rank {

			pass := false
			for i := king.Rank + 1; i < from.Rank; i++ {

				fmt.Println(board[king.File][i], i, king.Rank, from.Rank)
				if board[king.File][i].Type != Empty {
					pass = true

					break
				}
			}
			if !pass {
				for i := from.Rank + 1; i < 8; i++ {

					if board[king.File][i].Type != Empty {
						fmt.Print(board[king.File][i])
						if board[king.File][i].Type == Queen && board[king.File][i].Color != color {
							if king.File == to.File {
								return true, nil
							}
							return false, nil
						}
						if board[king.File][i].Type == Rook && board[king.File][i].Color != color {
							if king.File == to.File {
								return true, nil
							}
							return false, nil
						}
						break
					}
				}
			}

		} else {
			pass := false
			for i := king.Rank - 1; i > from.Rank; i-- {

				fmt.Println(board[king.File][i], i, king.Rank, from.Rank)
				if board[king.File][i].Type != Empty {
					pass = true
					fmt.Println(board[king.File][i])
					break
				}
			}
			if !pass {
				for i := from.Rank - 1; i >= 0; i-- {

					if board[king.File][i].Type != Empty {
						fmt.Print(board[king.File][i])
						if board[king.File][i].Type == Queen && board[king.File][i].Color != color {
							if king.File == to.File {
								return true, nil
							}
							return false, nil
						}
						if board[king.File][i].Type == Rook && board[king.File][i].Color != color {
							if king.File == to.File {
								return true, nil
							}
							return false, nil
						}
						break
					}
				}
			}
		}

	}

	return true, nil
}
