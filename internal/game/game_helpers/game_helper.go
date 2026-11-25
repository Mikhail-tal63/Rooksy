package gamehelpers

import (
	"fmt"
	"strconv"

	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_DTO"
	"github.com/Mikhail-Tal63/Rooksy/internal/game/game_errors"
)

func ParseFEN(fen string) (*Fen, error) {

	//    filling Board feild   *******************************

	var count, inner_c, after_break_counter int = 7, 0, 0
	var board [8][8]Piece

	for _, ch := range fen {
		after_break_counter++

		if ch == '/' {
			inner_c = 0
			count--
			continue
		} else if ch == ' ' {
			break
		}

		blank, peice, err := ParsePeice(ch)
		if err != nil {
			return nil, err
		}

		if peice == nil {
			peice = &Piece{
				Type:  Empty,
				Color: 0,
			}

			if inner_c+blank > 8 {
				return nil, fmt.Errorf("invalid FEN: rank exceeds 8 squares")
			}

			for j := inner_c; j < inner_c+blank; j++ {
				board[count][j] = *peice
			}

			inner_c += blank
		} else {
			if inner_c >= 8 {
				return nil, fmt.Errorf("invalid FEN: rank exceeds 8 squares")
			}

			board[count][inner_c] = *peice
			inner_c++
		}
	}

	// filling ActiveColor feild **************

	var active_color Colors

	if fen[after_break_counter] == 'w' {
		active_color = White
	} else {
		active_color = Black
	}

	//      filling Castling feild       88888888888
	after_break_counter += 2
	castle_string := ""
	for {
		if fen[after_break_counter] == ' ' {
			break
		} else {
			castle_string += string(fen[after_break_counter])
			after_break_counter++
		}
	}
	castling, err := ParseCastling(castle_string)
	if err != nil {
		return nil, err
	}
	//      filling EnPassant feild  ************

	after_break_counter++
	var enPassant string = ""
	if fen[after_break_counter] != '-' {
		enPassant += string(fen[after_break_counter])
		enPassant += string(fen[after_break_counter+1])
		after_break_counter++
	}

	//      filling fiftyrole feild
	after_break_counter += 2
	var ffty_string string = ""

	ffty_string += string(fen[after_break_counter])

	after_break_counter++

	if fen[after_break_counter] != ' ' {
		ffty_string += string(fen[after_break_counter])
		after_break_counter++
	}
	fifty_move, err := strconv.Atoi(ffty_string)
	if err != nil {
		return nil, game_errors.ErrInvalidHalfmoveClock
	}

	//      filling moves count feild

	after_break_counter++
	var moves_count_string string = ""

	moves_count_string += string(fen[after_break_counter])

	after_break_counter++

	if after_break_counter != len(fen) {
		moves_count_string += string(fen[after_break_counter])

	}
	moves_count, err := strconv.Atoi(moves_count_string)
	if err != nil {
		return nil, game_errors.ErrInvalidFullmoveNumber
	}

	finally_fen := &Fen{
		Board:         board,
		ActiveColor:   active_color,
		Castling:      castling,
		EnPassant:     enPassant,
		FiftyMove:     fifty_move,
		FullmoveCount: moves_count,
	}

	return finally_fen, nil
}

func ParsePeice(c rune) (int, *Piece, error) {

	var p Square
	switch c {
	case 'p', 'P':
		p = Pawn
	case 'r', 'R':
		p = Rook
	case 'n', 'N':
		p = Knight
	case 'b', 'B':
		p = Bishop
	case 'k', 'K':
		p = King
	case 'q', 'Q':
		p = Queen
	default:
		blank_squares, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, nil, fmt.Errorf("invalid FEN character: %c", c)
		}
		return blank_squares, nil, nil
	}
	co, err := ParseColor(c)
	if err != nil {
		return 0, nil, err
	}
	peice := &Piece{
		Type:  p,
		Color: co,
	}
	return 0, peice, nil
}

func ParseColor(c rune) (Colors, error) {
	if c >= 'A' && c <= 'Z' {
		return White, nil
	} else if c >= 'a' && c <= 'z' {
		return Black, nil
	}
	return 0, fmt.Errorf("Not valid piece %v", c)
}

func ParseCastling(s string) (CastlingRights, error) {
	var wk, wq, bk, bq bool = false, false, false, false
	castle := CastlingRights{
		WhiteKingSide:  wk,
		WhiteQueenSide: wq,
		BlackKingSide:  bk,
		BlackQueenSide: bq,
	}

	if s == "-" {
		return castle, nil
	}
	fmt.Print(s)
	for i := 0; i < len(s); i++ {

		switch s[i] {
		case 'K':
			wk = true
		case 'Q':
			wq = true
		case 'k':
			bk = true
		case 'q':
			bq = true
		}

	}
	castle = CastlingRights{
		WhiteKingSide:  wk,
		WhiteQueenSide: wq,
		BlackKingSide:  bk,
		BlackQueenSide: bq,
	}
	return castle, nil
}

func CreateFEN(fen *Fen) (string, error) {
	board := fen.Board
	fen_string := ""
	//      Board
	for i := 7; i >= 0; i-- {
		var Blank_count int = 0
		for j := 0; j < 8; j++ {
			if board[i][j].Type == 0 {
				Blank_count++
			} else {
				if Blank_count > 0 {
					fen_string += strconv.Itoa(Blank_count)
					Blank_count = 0
				}
				piece_str, err := CheckPieceType(board[i][j])
				if err != nil {
					return "", game_errors.ErrInvalidFENBoard
				}
				fen_string += piece_str
			}
		}
		if Blank_count > 0 {
			fen_string += strconv.Itoa(Blank_count)
		}
		if i != 0 {
			fen_string += "/"
		}
	}

	//   ActiveColor                       ************************

	if fen.ActiveColor == Black {
		fen_string += " b "
	} else if fen.ActiveColor == White {
		fen_string += " w "
	} else {
		return "", game_errors.ErrInvalidActiveColor
	}

	//  Castling *******************

	if fen.Castling.WhiteKingSide {
		fen_string += "K"
	}
	if fen.Castling.WhiteQueenSide {
		fen_string += "Q"
	}
	if fen.Castling.BlackKingSide {
		fen_string += "k"
	}
	if fen.Castling.BlackQueenSide {
		fen_string += "q"
	}
	if fen_string[len(fen_string)-1] == ' ' {
		fen_string += "-"
	}

	// EnPassant

	if fen.EnPassant == "" {
		fen_string += " -"
	} else {
		fen_string += " " + fen.EnPassant
	}

	//FiftyMove

	fen_string += " " + strconv.Itoa(fen.FiftyMove)

	//move count

	fen_string += " " + strconv.Itoa(fen.FullmoveCount)

	return fen_string, nil
}

func CheckPieceType(piece Piece) (string, error) {
	var color Colors = piece.Color
	var piece_string string
	if color == White {
		switch piece.Type {
		case Pawn:
			piece_string = "P"
		case Rook:
			piece_string = "R"
		case Knight:
			piece_string = "N"
		case Bishop:
			piece_string = "B"
		case King:
			piece_string = "K"
		case Queen:
			piece_string = "Q"
		default:
			return "", game_errors.ErrInvalidPiece
		}
	} else if color == Black {
		switch piece.Type {
		case Pawn:
			piece_string = "p"
		case Rook:
			piece_string = "r"
		case Knight:
			piece_string = "n"
		case Bishop:
			piece_string = "b"
		case King:
			piece_string = "k"
		case Queen:
			piece_string = "q"
		default:
			return "", game_errors.ErrInvalidPiece
		}
	} else {
		return "", game_errors.ErrInvalidPieceColor
	}
	return piece_string, nil
}

func GetKingPosition(fen *Fen, color Colors) (*Position, error) {
	board := fen.Board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j].Type == King && board[i][j].Color == color {
				pos := &Position{
					File: i,
					Rank: j,
				}
				return pos, nil
			}
		}
	}
	return nil, fmt.Errorf("King is not found")
}
