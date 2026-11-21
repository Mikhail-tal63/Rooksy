package gametests

import (
	"fmt"
	"testing"

	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func TestPinedPiece(t *testing.T) {

	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{
		{
			name:     "Vertical rook pin",
			from:     "4r2k/8/8/8/4R3/8/8/4K3 w - - 0 1",
			to:       "4r2k/8/8/8/8/4R3/8/4K3 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal rook pin",
			from:     "7k/8/8/8/r2R3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/8/r6K/8/8/3R4 w - - 0 1",
			expected: true,
		},
		{
			name:     "No pin - king moves",
			from:     "4r2k/8/8/8/4R3/8/8/3K4 w - - 0 1",
			to:       "4r2k/8/8/8/4R3/8/8/2K5 w - - 0 1",
			expected: false,
		},
		{
			name:     "Vertical queen pin",
			from:     "3q3k/8/8/8/3R4/8/8/3K4 w - - 0 1",
			to:       "3q3k/8/8/8/8/3R4/8/3K4 w - - 0 1",
			expected: true,
		},
	}

	for _, tc := range tests {
		pars1, err := ParseFEN(tc.from)
		if err != nil {
			panic(err)
		}

		pars2, err := ParseFEN(tc.to)
		if err != nil {
			panic(err)
		}

		moves, err := CalcolateMoves(pars1, pars2)
		if err != nil {
			panic(err)
		}

		result, err := CheckPinedPeiceHorVir(pars1, *moves)
		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"%-25s Expected: %-5v Got: %-5v\n",
			tc.name, tc.expected, result,
		)
	}

}
