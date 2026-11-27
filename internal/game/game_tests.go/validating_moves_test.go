package gametests

import (
	"testing"

	game "github.com/Mikhail-Tal63/Rooksy/internal/game/game_engine/game_validate"
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func TestValidateknightMove(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{{
		name:     "Knight d4-e6",
		from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
		to:       "7k/8/4N3/8/8/8/8/K7 w - - 0 1",
		expected: true,
	},
		{
			name:     "Knight d4-f5",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/5N2/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight d4-f3",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/8/5N2/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight d4-e2",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/8/8/4N3/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight d4-c2",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/8/8/2N5/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight d4-b3",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/8/1N6/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight d4-b5",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/1N6/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight d4-c6",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/2N5/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight illegal vertical d4-d6",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/3N4/8/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Knight illegal horizontal d4-f4",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/5N2/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Knight illegal diagonal d4-e5",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/4N3/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Knight moves too far d4-g5",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/6N1/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Knight illegal adjacent d4-d5",
			from:     "7k/8/8/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/3N4/8/8/8/K7 w - - 0 1",
			expected: false,
		}, {
			name:     "Knight jumps over friendly and enemy pieces",
			from:     "7k/8/8/3PP3/3Npp2/8/8/K7 w - - 0 1",
			to:       "7k/8/8/3PPN2/4pp2/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight captures enemy pawn",
			from:     "7k/8/8/5p2/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/5N2/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight captures enemy rook",
			from:     "7k/8/2r5/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/2N5/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight lands on friendly pawn",
			from:     "7k/8/8/5P2/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/8/5N2/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Knight lands on friendly bishop",
			from:     "7k/8/2B5/8/3N4/8/8/K7 w - - 0 1",
			to:       "7k/8/2N5/8/8/8/8/K7 w - - 0 1",
			expected: false,
		}, {
			name:     "Knight moves from corner a8-b6",
			from:     "N6k/8/8/8/8/8/8/K7 w - - 0 1",
			to:       "7k/8/1N6/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Knight moves to corner b6-a8",
			from:     "7k/8/1N6/8/8/8/8/K7 w - - 0 1",
			to:       "N6k/8/8/8/8/8/8/K7 w - - 0 1",
			expected: true,
		}, {
			name:     "Black knight legal d5-f4",
			from:     "7k/8/8/3n4/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/5n2/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black knight captures white pawn",
			from:     "7k/8/8/3n4/8/2P5/8/K7 b - - 0 1",
			to:       "7k/8/8/8/8/2n5/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black knight lands on friendly pawn",
			from:     "7k/8/8/3n4/8/2p5/8/K7 b - - 0 1",
			to:       "7k/8/8/8/8/2n5/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black knight illegal vertical d5-d3",
			from:     "7k/8/8/3n4/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/8/3n4/8/K7 b - - 0 1",
			expected: false,
		}}
	for _, test := range tests {
		parseFrom, err := ParseFEN(test.from)
		if err != nil {
			panic(err)
		}
		parseTo, err := ParseFEN(test.to)
		if err != nil {
			panic(err)
		}

		valid, err := game.ValidateKnightMove(parseFrom, parseTo)
		if err != nil {
			panic(err)
		}
		if valid != test.expected {
			t.Errorf(
				"%s: expected %v, got %v",
				test.name, test.expected, valid,
			)
		}
	}
}
