package gametests

import (
	"testing"

	game "github.com/Mikhail-Tal63/Rooksy/internal/game/game_engine/game_validate"
	. "github.com/Mikhail-Tal63/Rooksy/internal/game/game_helpers"
)

func TestVerticalHorizontalPin(t *testing.T) {

	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{{
		name:     "Vertical rook pin from above - knight",
		from:     "4r2k/8/8/8/4N3/8/8/4K3 w - - 0 1",
		to:       "4r2k/8/8/2N5/8/8/8/4K3 w - - 0 1",
		expected: false,
	},
		{
			name:     "Vertical rook pin from below - bishop",
			from:     "3K4/8/8/3B4/8/8/8/3r3k w - - 0 1",
			to:       "3K4/8/8/8/4B3/8/8/3r3k w - - 0 1",
			expected: false,
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
			expected: false,
		},
		{
			name:     "Vertical pin - black knight",
			from:     "4k3/8/8/4n3/8/8/8/K3R3 b - - 0 1",
			to:       "4k3/8/8/8/2n5/8/8/K3R3 b - - 0 1",
			expected: false,
		},
		{
			name:     "Vertical pin - black bishop by white queen",
			from:     "2k5/8/8/2b5/8/8/8/2Q4K b - - 0 1",
			to:       "2k5/8/8/8/3b4/8/8/2Q4K b - - 0 1",
			expected: false,
		},
		{
			name:     "Vertical pin on edge file",
			from:     "r6k/8/8/8/R7/8/8/K7 w - - 0 1",
			to:       "r6k/8/8/8/1R6/8/8/K7 w - - 0 1",
			expected: false,
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
			expected: false,
		},
		{
			name:     "Horizontal rook pin from right - bishop",
			from:     "7k/8/8/8/K2B3r/8/8/8 w - - 0 1",
			to:       "7k/8/8/4B3/K6r/8/8/8 w - - 0 1",
			expected: false,
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
			expected: false,
		},
		{
			name:     "Horizontal pin - black knight",
			from:     "8/8/8/R2n3k/8/8/8/7K b - - 0 1",
			to:       "8/8/5n2/R6k/8/8/8/7K b - - 0 1",
			expected: false,
		},
		{
			name:     "Horizontal pin - black rook by white queen",
			from:     "8/8/8/k2r3Q/8/8/8/K7 b - - 0 1",
			to:       "8/8/8/k6Q/3r4/8/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Horizontal pin - pawn",
			from:     "7k/8/8/8/r2P3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/3P4/r6K/8/8/8 w - - 0 1",
			expected: false,
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
			expected: true,
		},
		{
			name:     "No horizontal pin - two friendly blockers",
			from:     "7k/8/8/8/r1B1R2K/8/8/8 w - - 0 1",
			to:       "7k/8/8/4R3/r1B4K/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "No vertical pin - attacker is bishop",
			from:     "4b2k/8/8/8/4R3/8/8/4K3 w - - 0 1",
			to:       "4b2k/8/8/8/5R2/8/8/4K3 w - - 0 1",
			expected: true,
		},
		{
			name:     "No horizontal pin - attacker is bishop",
			from:     "7k/8/8/8/b2R3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/3R4/b6K/8/8/8 w - - 0 1",
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

		moves, err := LocateMove(pars1, pars2)
		if err != nil {
			panic(err)
		}

		result, err := CheckPinedPeiceHorVir(pars1, *moves)
		if err != nil {
			panic(err)
		}

		if result != tc.expected {
			t.Errorf(
				"%s: expected %v, got %v",
				tc.name, tc.expected, result,
			)
		}
	}

}

func TestDiagnlePin(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{{
		name:     "Diagonal bishop pin NE",
		from:     "k6b/8/8/8/3R4/8/8/K7 w - - 0 1",
		to:       "k6b/8/8/8/4R3/8/8/K7 w - - 0 1",
		expected: false,
	},
		{
			name:     "Diagonal bishop pin NW",
			from:     "b6k/8/8/8/4R3/8/8/7K w - - 0 1",
			to:       "b6k/8/8/8/3R4/8/8/7K w - - 0 1",
			expected: false,
		},
		{
			name:     "Diagonal bishop pin SE",
			from:     "K6k/8/8/3R4/8/8/8/7b w - - 0 1",
			to:       "K6k/8/8/4R3/8/8/8/7b w - - 0 1",
			expected: false,
		},
		{
			name:     "Diagonal bishop pin SW",
			from:     "k6K/8/8/4R3/8/8/8/b7 w - - 0 1",
			to:       "k6K/8/8/3R4/8/8/8/b7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black diagonal bishop pin NE",
			from:     "K6B/8/8/8/3r4/8/8/k7 b - - 0 1",
			to:       "K6B/8/8/8/4r3/8/8/k7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black diagonal bishop pin NW",
			from:     "B6K/8/8/8/4r3/8/8/7k b - - 0 1",
			to:       "B6K/8/8/8/3r4/8/8/7k b - - 0 1",
			expected: false,
		},
		{
			name:     "Black diagonal queen pin SE",
			from:     "k6K/8/8/3r4/8/8/8/7Q b - - 0 1",
			to:       "k6K/8/8/4r3/8/8/8/7Q b - - 0 1",
			expected: false,
		},
		{
			name:     "Black diagonal queen pin SW",
			from:     "K6k/8/8/4r3/8/8/8/Q7 b - - 0 1",
			to:       "K6k/8/8/3r4/8/8/8/Q7 b - - 0 1",
			expected: false,
		},
		{
			name:     "No diagonal pin NE - two blockers",
			from:     "k6b/8/8/4N3/3R4/8/8/K7 w - - 0 1",
			to:       "k6b/8/8/4N3/4R3/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin NW - two blockers",
			from:     "b6k/8/8/3N4/4R3/8/8/7K w - - 0 1",
			to:       "b6k/8/8/3N4/3R4/8/8/7K w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin SE - two blockers",
			from:     "K6k/8/8/3R4/4N3/8/8/7b w - - 0 1",
			to:       "K6k/8/8/4R3/4N3/8/8/7b w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin SW - two blockers",
			from:     "k6K/8/8/4R3/3N4/8/8/b7 w - - 0 1",
			to:       "k6K/8/8/3R4/3N4/8/8/b7 w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin - attacker is rook",
			from:     "k6r/8/8/8/3R4/8/8/K7 w - - 0 1",
			to:       "k6r/8/8/8/4R3/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin - friendly bishop",
			from:     "k6B/8/8/8/3R4/8/8/K7 w - - 0 1",
			to:       "k6B/8/8/8/4R3/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin - piece outside pin line",
			from:     "k6b/8/8/4P3/4R3/8/8/K7 w - - 0 1",
			to:       "k6b/8/8/4P3/5R2/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "No diagonal pin - enemy blocker",
			from:     "k6b/8/8/4n3/3R4/8/8/K7 w - - 0 1",
			to:       "k6b/8/8/4n3/4R3/8/8/K7 w - - 0 1",
			expected: true,
		},
	}

	for _, test := range tests {
		parseFrom, err := ParseFEN(test.from)
		if err != nil {
			panic(err)
		}
		parseTo, err := ParseFEN(test.to)
		if err != nil {
			panic(err)
		}
		move, err := LocateMove(parseFrom, parseTo)
		if err != nil {
			panic(err)
		}
		pin, err := CheckPinedDiagnly(parseFrom, *move)
		if err != nil {
			panic(err)
		}
		if pin != test.expected {
			t.Errorf(
				"%s: expected %v, got %v",
				test.name, test.expected, pin,
			)
		}
	}
}

func TestValidateBishopMove(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{
		{
			name:     "NE bishop clear",
			from:     "8/8/8/7k/K7/8/1B6/8 w - - 0 1",
			to:       "8/8/5B2/7k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "NE queen clear",
			from:     "8/8/8/7k/K7/8/1Q6/8 w - - 0 1",
			to:       "8/8/5Q2/7k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "NE bishop friendly blocker",
			from:     "8/8/8/7k/K2P4/8/1B6/8 w - - 0 1",
			to:       "8/8/5B2/7k/K2P4/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "NE bishop enemy blocker",
			from:     "8/8/8/7k/K2p4/8/1B6/8 w - - 0 1",
			to:       "8/8/5B2/7k/K2p4/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "NW bishop clear",
			from:     "8/8/8/7k/K7/8/6B1/8 w - - 0 1",
			to:       "8/8/2B5/7k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "NW queen clear",
			from:     "8/8/8/7k/K7/8/6Q1/8 w - - 0 1",
			to:       "8/8/2Q5/7k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "NW bishop friendly blocker",
			from:     "8/8/8/7k/K3P3/8/6B1/8 w - - 0 1",
			to:       "8/8/2B5/7k/K3P3/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "NW bishop enemy blocker",
			from:     "8/8/8/7k/K3p3/8/6B1/8 w - - 0 1",
			to:       "8/8/2B5/7k/K3p3/8/8/8 w - - 0 1",
			expected: false,
		}, {
			name:     "SE bishop clear",
			from:     "8/1B6/8/7k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/7k/K7/5B2/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SE queen clear",
			from:     "8/1Q6/8/7k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/7k/K7/5Q2/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SE bishop friendly blocker",
			from:     "8/1B6/8/3P3k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/3P3k/K7/5B2/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "SE bishop enemy blocker",
			from:     "8/1B6/8/3p3k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/3p3k/K7/5B2/8/8 w - - 0 1",
			expected: false,
		}, {
			name:     "SW bishop clear",
			from:     "8/6B1/8/7k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/7k/K7/2B5/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SW queen clear",
			from:     "8/6Q1/8/7k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/7k/K7/2Q5/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SW bishop friendly blocker",
			from:     "8/6B1/8/4P2k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/4P2k/K7/2B5/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "SW bishop enemy blocker",
			from:     "8/6B1/8/4p2k/K7/8/8/8 w - - 0 1",
			to:       "8/8/8/4p2k/K7/2B5/8/8 w - - 0 1",
			expected: false,
		}, {
			name:     "Adjacent diagonal",
			from:     "8/8/8/7k/K2B4/8/8/8 w - - 0 1",
			to:       "8/8/8/4B2k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Enemy on destination",
			from:     "8/8/5n2/7k/K7/8/1B6/8 w - - 0 1",
			to:       "8/8/5B2/7k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Blocker immediately after source",
			from:     "8/8/8/7k/K7/2p5/1B6/8 w - - 0 1",
			to:       "8/8/5B2/7k/K7/2p5/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "Blocker immediately before destination",
			from:     "8/8/8/4P2k/K7/8/1B6/8 w - - 0 1",
			to:       "8/8/5B2/4P2k/K7/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "Piece beyond destination",
			from:     "8/8/8/4p2k/K7/8/1B6/8 w - - 0 1",
			to:       "8/8/8/4p2k/K2B4/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Piece outside path",
			from:     "8/8/8/3P3k/K7/8/1B6/8 w - - 0 1",
			to:       "8/8/5B2/3P3k/K7/8/8/8 w - - 0 1",
			expected: true,
		},
	}
	for _, test := range tests {
		parseFrom, err := ParseFEN(test.from)
		if err != nil {
			panic(err)
		}
		parseTo, err := ParseFEN(test.to)
		if err != nil {
			panic(err)
		}
		move, err := LocateMove(parseFrom, parseTo)
		if err != nil {
			panic(err)
		}
		pin, err := ValidateDiagnalMoves(*move, *parseFrom)
		if err != nil {
			panic(err)
		}
		if pin != test.expected {
			t.Errorf(
				"%s: expected %v, got %v",
				test.name, test.expected, pin,
			)
		}
	}
}

func TestValidateQueenMove(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{{
		name:     "Queen North d4-d7",
		from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
		to:       "8/3Q3k/8/8/8/8/8/K7 w - - 0 1",
		expected: true,
	},
		{
			name:     "Queen South d4-d1",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/8/8/8/8/K2Q4 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen East d4-g4",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/8/6Q1/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen West d4-a4",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/8/Q7/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen NE d4-g7",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6Qk/8/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen NW d4-a7",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/Q6k/8/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen SE d4-g1",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/8/8/8/8/K5Q1 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen SW d4-b2",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/8/8/8/1Q6/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen vertical friendly blocker",
			from:     "8/7k/8/3P4/3Q4/8/8/K7 w - - 0 1",
			to:       "8/3Q3k/8/3P4/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Queen diagonal enemy blocker",
			from:     "8/7k/8/4p3/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6Qk/8/4p3/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Queen horizontal enemy blocker",
			from:     "8/7k/8/8/3Q1n2/8/8/K7 w - - 0 1",
			to:       "8/7k/8/8/5nQ1/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Queen captures enemy rook",
			from:     "8/6rk/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6Qk/8/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen illegal knight movement d4-f5",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/5Q2/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Queen invalid diagonal d4-g6",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/6Q1/8/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Queen adjacent diagonal",
			from:     "8/7k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/7k/8/4Q3/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen piece beyond destination",
			from:     "8/3p3k/8/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/3p3k/3Q4/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pinned queen moves along pin",
			from:     "4r3/8/7k/8/4Q3/8/8/4K3 w - - 0 1",
			to:       "4r3/8/4Q2k/8/8/8/8/4K3 w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pinned queen moves horizontally",
			from:     "4r3/8/7k/8/4Q3/8/8/4K3 w - - 0 1",
			to:       "4r3/8/7k/8/5Q2/8/8/4K3 w - - 0 1",
			expected: false,
		},
		{
			name:     "Vertical pinned queen captures attacker",
			from:     "4r3/8/7k/8/4Q3/8/8/4K3 w - - 0 1",
			to:       "4Q3/8/7k/8/8/8/8/4K3 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal pinned queen moves along pin",
			from:     "8/8/7k/8/K2Q3r/8/8/8 w - - 0 1",
			to:       "8/8/7k/8/K4Q1r/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "Horizontal pinned queen moves vertically",
			from:     "8/8/7k/8/K2Q3r/8/8/8 w - - 0 1",
			to:       "8/8/7k/3Q4/K6r/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "Horizontal pinned queen captures attacker",
			from:     "8/8/7k/8/K2Q3r/8/8/8 w - - 0 1",
			to:       "8/8/7k/8/K6Q/8/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "NE pinned queen moves along pin",
			from:     "8/6b1/7k/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6b1/7k/4Q3/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "NE pinned queen breaks pin",
			from:     "8/6b1/7k/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6b1/7k/3Q4/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "NE pinned queen captures bishop",
			from:     "8/6b1/7k/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6Q1/7k/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "NW pinned queen moves along pin",
			from:     "8/1b6/k7/8/4Q3/8/8/7K w - - 0 1",
			to:       "8/1b6/k7/3Q4/8/8/8/7K w - - 0 1",
			expected: true,
		},
		{
			name:     "NW pinned queen breaks pin",
			from:     "8/1b6/k7/8/4Q3/8/8/7K w - - 0 1",
			to:       "8/1b6/k7/4Q3/8/8/8/7K w - - 0 1",
			expected: false,
		},
		{
			name:     "NW pinned queen captures bishop",
			from:     "8/1b6/k7/8/4Q3/8/8/7K w - - 0 1",
			to:       "8/1Q6/k7/8/8/8/8/7K w - - 0 1",
			expected: true,
		},
		{
			name:     "SE pinned queen moves along pin",
			from:     "K7/8/8/3Q4/8/7k/6b1/8 w - - 0 1",
			to:       "K7/8/8/8/4Q3/7k/6b1/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SE pinned queen breaks pin",
			from:     "K7/8/8/3Q4/8/7k/6b1/8 w - - 0 1",
			to:       "K7/8/8/8/3Q4/7k/6b1/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "SE pinned queen captures bishop",
			from:     "K7/8/8/3Q4/8/7k/6b1/8 w - - 0 1",
			to:       "K7/8/8/8/8/7k/6Q1/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SW pinned queen moves along pin",
			from:     "7K/8/8/4Q3/8/k7/1b6/8 w - - 0 1",
			to:       "7K/8/8/8/3Q4/k7/1b6/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "SW pinned queen breaks pin",
			from:     "7K/8/8/4Q3/8/k7/1b6/8 w - - 0 1",
			to:       "7K/8/8/8/4Q3/k7/1b6/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "SW pinned queen captures bishop",
			from:     "7K/8/8/4Q3/8/k7/1b6/8 w - - 0 1",
			to:       "7K/8/8/8/8/k7/1Q6/8 w - - 0 1",
			expected: true,
		}, {
			name:     "Black pinned queen moves toward king",
			from:     "4k3/8/8/4q3/8/7K/8/4R3 b - - 0 1",
			to:       "4k3/4q3/8/8/8/7K/8/4R3 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black pinned queen breaks pin",
			from:     "4k3/8/8/4q3/8/7K/8/4R3 b - - 0 1",
			to:       "4k3/8/8/5q2/8/7K/8/4R3 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black pinned queen captures white rook",
			from:     "4k3/8/8/4q3/8/7K/8/4R3 b - - 0 1",
			to:       "4k3/8/8/8/8/7K/8/4q3 b - - 0 1",
			expected: true,
		}, {
			name:     "Queen pinned by queen moves toward king",
			from:     "8/6q1/7k/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6q1/7k/8/8/2Q5/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen pinned by queen breaks pin",
			from:     "8/6q1/7k/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6q1/7k/8/4Q3/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Queen captures pinning queen",
			from:     "8/6q1/7k/8/3Q4/8/8/K7 w - - 0 1",
			to:       "8/6Q1/7k/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Queen blocks existing check",
			from:     "4r3/7k/8/8/3Q4/8/8/4K3 w - - 0 1",
			to:       "4r3/7k/8/8/4Q3/8/8/4K3 w - - 0 1",
			expected: true,
		},
	}
	for _, test := range tests {
		parseFrom, err := ParseFEN(test.from)
		if err != nil {
			panic(err)
		}
		parseTo, err := ParseFEN(test.to)
		if err != nil {
			panic(err)
		}

		pin, err := game.ValidateQueenMove(parseFrom, parseTo)
		if err != nil {
			panic(err)
		}
		if pin != test.expected {
			t.Errorf(
				"%s: expected %v, got %v",
				test.name, test.expected, pin,
			)
		}
	}
}
