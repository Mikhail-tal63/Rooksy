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
	}{{
		name:     "Vertical rook pin from above - knight",
		from:     "4r2k/8/8/8/4N3/8/8/4K3 w - - 0 1",
		to:       "4r2k/8/8/2N5/8/8/8/4K3 w - - 0 1",
		expected: true,
	},
		{
			name:     "Vertical rook pin from below - bishop",
			from:     "3K4/8/8/3B4/8/8/8/3r3k w - - 0 1",
			to:       "3K4/8/8/8/4B3/8/8/3r3k w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical queen pin - pawn moves along pin",
			from:     "4q2k/8/8/8/8/8/4P3/4K3 w - - 0 1",
			to:       "4q2k/8/8/8/8/4P3/8/4K3 w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical queen pin from below - rook",
			from:     "2K5/8/8/2R5/8/8/8/2q4k w - - 0 1",
			to:       "2K5/8/8/5R2/8/8/8/2q4k w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pin - black knight",
			from:     "4k3/8/8/4n3/8/8/8/K3R3 b - - 0 1",
			to:       "4k3/8/8/8/2n5/8/8/K3R3 b - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pin - black bishop by white queen",
			from:     "2k5/8/8/2b5/8/8/8/2Q4K b - - 0 1",
			to:       "2k5/8/8/8/3b4/8/8/2Q4K b - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pin on edge file",
			from:     "r6k/8/8/8/R7/8/8/K7 w - - 0 1",
			to:       "r6k/8/8/8/1R6/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pin - rook captures pinning rook",
			from:     "4r2k/8/8/8/4R3/8/8/4K3 w - - 0 1",
			to:       "4R2k/8/8/8/8/8/8/4K3 w - - 0 1",
			expected: true,
		}, {
			name:     "Horizontal rook pin from left - knight",
			from:     "7k/8/8/8/r2N3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/5N2/r6K/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal rook pin from right - bishop",
			from:     "7k/8/8/8/K2B3r/8/8/8 w - - 0 1",
			to:       "7k/8/8/4B3/K6r/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal queen pin from left",
			from:     "7k/8/8/8/q2R3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/8/q3R2K/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal queen pin from right",
			from:     "7k/8/8/8/K2R3q/8/8/8 w - - 0 1",
			to:       "7k/8/8/3R4/K6q/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal pin - black knight",
			from:     "8/8/8/R2n3k/8/8/8/7K b - - 0 1",
			to:       "8/8/5n2/R6k/8/8/8/7K b - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal pin - black rook by white queen",
			from:     "8/8/8/k2r3Q/8/8/8/K7 b - - 0 1",
			to:       "8/8/8/k6Q/3r4/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal pin - pawn",
			from:     "7k/8/8/8/r2P3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/3P4/r6K/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal pin - rook captures pinning queen",
			from:     "7k/8/8/8/q2R3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/8/R6K/8/8/8 w - - 0 1",
			expected: true,
		}, {
			name:     "No vertical pin - two friendly blockers",
			from:     "4r2k/8/8/4B3/4R3/8/8/4K3 w - - 0 1",
			to:       "4r2k/8/8/4B3/5R2/8/8/4K3 w - - 0 1",
			expected: false,
		},
		{
			name:     "No horizontal pin - two friendly blockers",
			from:     "7k/8/8/8/r1B1R2K/8/8/8 w - - 0 1",
			to:       "7k/8/8/4R3/r1B4K/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "No vertical pin - attacker is bishop",
			from:     "4b2k/8/8/8/4R3/8/8/4K3 w - - 0 1",
			to:       "4b2k/8/8/8/5R2/8/8/4K3 w - - 0 1",
			expected: false,
		},
		{
			name:     "No horizontal pin - attacker is bishop",
			from:     "7k/8/8/8/b2R3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/3R4/b6K/8/8/8 w - - 0 1",
			expected: false,
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

		moves, err := LocateMove(pars1, pars2)
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
