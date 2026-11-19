package game

import (
	"fmt"
	"strconv"
)

func ParseFEN(fen string) (*Fen, error) {
	//    filling Board feild   *******************************
	var count, inner_c, after_break_counter int = 0, 0, 0
	var board [8][8]Piece
	for _, ch := range fen {
		after_break_counter++
		if ch == '/' {
			inner_c = 0
			count++

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
	fmt.Print(string(fen[after_break_counter]))
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
		return nil, err
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
		return nil, err
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
