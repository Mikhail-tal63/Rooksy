package gamehelpers

import (
	"fmt"
	"math"

	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
)

/*
P_somthing is refrence for pined peice
*/

// returns the move the have been made
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

	hv, err := CheckPinedPeiceHorVir(fen, move)
	if err != nil {
		return false, err
	}
	if !hv {
		return false, nil
	}
	return true, nil
}
func CheckPinedPeiceHorVir(fen *Fen, move Move) (bool, error) {
	from := move.From
	to := move.To
	board := fen.Board
	color := fen.Board[from.File][from.Rank].Color
	king, err := GetKingPosition(fen, color)
	if err != nil {
		return false, fmt.Errorf("could not find king : validate_helper")
	}

	if from.File == king.File {

		ra, err := P_CheckRank(board, from, color, *king, to)
		if err != nil {
			return false, err
		}
		if !ra {
			return false, nil
		}
	}
	if from.Rank == king.Rank {
		fi, err := P_CheckFile(board, from, color, *king, to)
		if err != nil {
			return false, err
		}
		if !fi {
			return false, nil
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
func P_CheckFile(board [8][8]Piece, from Position, color Colors, king Position, to Position) (bool, error) {

	if from.File > king.File {

		pass := false
		for i := king.File + 1; i < from.File; i++ {

			fmt.Println(board[i][king.Rank], i, king.File, from.File)
			if board[i][king.Rank].Type != Empty {
				pass = true

				break
			}
		}
		if !pass {
			for i := from.File + 1; i < 8; i++ {

				if board[i][king.Rank].Type != Empty {
					fmt.Print(board[i][king.Rank])
					if board[i][king.Rank].Type == Queen && board[i][king.Rank].Color != color {
						if king.Rank == to.Rank {
							return true, nil
						}
						fmt.Print(board[i][king.Rank])
						return false, nil
					}
					if board[i][king.Rank].Type == Rook && board[i][king.Rank].Color != color {
						if king.Rank == to.Rank {
							return true, nil
						}
						fmt.Print(board[i][king.Rank])
						return false, nil
					}
					break
				}
			}
		}

	} else {
		pass := false
		for i := king.File - 1; i > from.File; i-- {

			fmt.Println(board[i][king.Rank], i, king.File, from.File)
			if board[i][king.Rank].Type != Empty {
				pass = true
				fmt.Println(board[i][king.Rank])
				break
			}
		}
		if !pass {
			for i := from.File - 1; i >= 0; i-- {

				if board[i][king.Rank].Type != Empty {
					fmt.Print(board[i][king.Rank])
					if board[i][king.Rank].Type == Queen && board[i][king.Rank].Color != color {
						if king.Rank == to.Rank {
							return true, nil
						}
						return false, nil
					}
					if board[i][king.Rank].Type == Rook && board[i][king.Rank].Color != color {
						if king.Rank == to.Rank {
							return true, nil
						}
						return false, nil
					}
					break
				}
			}
		}
	}

	return true, nil
}
func P_CheckUpperRightDiagnal(board [8][8]Piece, from Position, color Colors, king Position, to Position) (bool, error) {
	l := king.Rank + 1
	r := king.File + 1
	fmt.Println(board[r][l], l, r, from.Rank, from.File, "hellorrr")
	for l < from.Rank && r < from.File {

		if board[r][l].Type != Empty {

			return true, nil
		}
		r++
		l++
	}
	l = from.Rank + 1
	r = from.File + 1
	for l < 8 && r < 8 {
		fmt.Println(board[r][l], l, r, king.File, from.File)
		if board[r][l].Type != Empty {

			if board[r][l].Type == Queen && board[r][l].Color != color {
				if to.File-to.Rank != from.File-from.Rank {
					return false, nil
				}
				return true, nil

			}
			if board[r][l].Type == Bishop && board[r][l].Color != color {
				if to.File-to.Rank != from.File-from.Rank {
					return false, nil
				}
				return true, nil

				//if you suptact from-from and to -to and got same answear you good else bad
			}
		}
		r++
		l++
	}
	return true, nil
}
func P_CheckUpperLeftDiagnal(board [8][8]Piece, from Position, color Colors, king Position, to Position) (bool, error) {
	l := king.Rank - 1
	r := king.File + 1
	fmt.Println(board[r][l], l, r, from.Rank, from.File, "hellodsgfrrr")
	for l > from.Rank && r < from.File {

		if board[r][l].Type != Empty {

			return true, nil
		}
		r++
		l--
	}
	l = from.Rank - 1
	r = from.File + 1
	for l >= 0 && r < 8 {
		fmt.Println(board[r][l], l, r, from.Rank, from.File)
		if board[r][l].Type != Empty {

			if board[r][l].Type == Queen && board[r][l].Color != color {
				if math.Abs(float64(to.File)-float64(from.Rank)) != math.Abs(float64(to.Rank)-float64(from.File)) {
					return false, nil
				}
				return true, nil

			}
			if board[r][l].Type == Bishop && board[r][l].Color != color {
				fmt.Println(board[r][l], to.Rank, to.File, from.Rank, from.File)
				if math.Abs(float64(to.File)-float64(from.Rank)) != math.Abs(float64(to.Rank)-float64(from.File)) {
					return false, nil
				}
				return true, nil

				//if you suptact from-from and to -to and got same answear you good else bad
			}
		}
		r++
		l--
	}
	return true, nil
}
func P_CheckLowerRightDiagnal(board [8][8]Piece, from Position, color Colors, king Position, to Position) (bool, error) {
	l := king.Rank + 1
	r := king.File - 1
	for l < from.Rank && r > from.File {

		if board[r][l].Type != Empty {

			return true, nil
		}
		r--
		l++
	}
	l = from.Rank + 1
	r = from.File - 1
	for l < 8 && r >= 0 {

		if board[r][l].Type != Empty {

			if board[r][l].Type == Queen && board[r][l].Color != color {
				if from.File-from.Rank == to.File-to.Rank {
					return false, nil
				}
				if from.File == to.File || to.Rank == from.Rank {
					return false, nil
				}
				return true, nil

			}
			if board[r][l].Type == Bishop && board[r][l].Color != color {
				fmt.Println(board[r][l], to.Rank, to.File, from.Rank, from.File, "hi")
				if from.File-from.Rank == to.File-to.Rank {
					return false, nil
				}
				if from.File == to.File || to.Rank == from.Rank {
					return false, nil
				}
				return true, nil

				//if you suptact from-from and to -to and got same answear you good else bad
			}
		}
		r--
		l++
	}
	return true, nil
}
