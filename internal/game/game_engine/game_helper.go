package game

import (
	"fmt"
)

func ParseFEN(fen string) (*Fen, error) {
	for _, ch := range fen {
		ch
	}
}

func ParsePeice(c rune) (int, *Piece, error) {
	co, err := ParseColor(c)
	if err != nil {
		return 0, nil, err
	}
	var blank_squares int
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
		blank_squares = int(c)
		return blank_squares, nil, nil
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
	return 0, fmt.Errorf("Not valid piece")
}
