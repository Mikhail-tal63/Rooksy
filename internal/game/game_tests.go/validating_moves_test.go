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
func TestValidatePawnMove(t *testing.T) {
	tests := []struct {
		name     string
		from     string
		to       string
		expected bool
	}{{
		name:     "White pawn single push e2-e3",
		from:     "7k/8/8/8/8/8/4P3/K7 w - - 0 1",
		to:       "7k/8/8/8/8/4P3/8/K7 w - - 0 1",
		expected: true,
	},
		{
			name:     "White pawn double push e2-e4",
			from:     "7k/8/8/8/8/8/4P3/K7 w - - 0 1",
			to:       "7k/8/8/8/4P3/8/8/K7 w - e3 0 1",
			expected: true,
		},
		{
			name:     "White pawn single push from middle e4-e5",
			from:     "7k/8/8/8/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/4P3/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "White pawn illegal double push e3-e5",
			from:     "7k/8/8/8/8/4P3/8/K7 w - - 0 1",
			to:       "7k/8/8/4P3/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn illegal backward e4-e3",
			from:     "7k/8/8/8/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/8/4P3/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn illegal horizontal e4-f4",
			from:     "7k/8/8/8/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/8/5P2/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn illegal triple push e2-e5",
			from:     "7k/8/8/8/8/8/4P3/K7 w - - 0 1",
			to:       "7k/8/8/4P3/8/8/8/K7 w - - 0 1",
			expected: false,
		}, {
			name:     "Black pawn single push e7-e6",
			from:     "7k/4p3/8/8/8/8/8/K7 b - - 0 1",
			to:       "7k/8/4p3/8/8/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black pawn double push e7-e5",
			from:     "7k/4p3/8/8/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/4p3/8/8/8/K7 b - e6 0 1",
			expected: true,
		},
		{
			name:     "Black pawn single push e5-e4",
			from:     "7k/8/8/4p3/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/4p3/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black pawn illegal double push e6-e4",
			from:     "7k/8/4p3/8/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/4p3/8/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black pawn illegal backward e5-e6",
			from:     "7k/8/8/4p3/8/8/8/K7 b - - 0 1",
			to:       "7k/8/4p3/8/8/8/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black pawn illegal horizontal e5-d5",
			from:     "7k/8/8/4p3/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/3p4/8/8/8/K7 b - - 0 1",
			expected: false,
		}, {
			name:     "White pawn forward friendly blocker",
			from:     "7k/8/8/8/8/4N3/4P3/K7 w - - 0 1",
			to:       "7k/8/8/8/8/4P3/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn forward enemy blocker",
			from:     "7k/8/8/8/8/4n3/4P3/K7 w - - 0 1",
			to:       "7k/8/8/8/8/4P3/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn double push friendly intermediate blocker",
			from:     "7k/8/8/8/8/4N3/4P3/K7 w - - 0 1",
			to:       "7k/8/8/8/4P3/4N3/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn double push enemy intermediate blocker",
			from:     "7k/8/8/8/8/4n3/4P3/K7 w - - 0 1",
			to:       "7k/8/8/8/4P3/4n3/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn double push occupied destination",
			from:     "7k/8/8/8/4n3/8/4P3/K7 w - - 0 1",
			to:       "7k/8/8/8/4P3/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black pawn double push intermediate blocker",
			from:     "7k/4p3/4N3/8/8/8/8/K7 b - - 0 1",
			to:       "7k/8/4N3/4p3/8/8/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black pawn forward enemy blocker",
			from:     "7k/8/8/4p3/4P3/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/4p3/8/8/K7 b - - 0 1",
			expected: false,
		}, {
			name:     "White pawn captures left e4-d5",
			from:     "7k/8/8/3p4/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/3P4/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "White pawn captures right e4-f5",
			from:     "7k/8/8/5n2/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/5P2/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "White pawn diagonal empty square",
			from:     "7k/8/8/8/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/5P2/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn captures friendly piece",
			from:     "7k/8/8/3N4/4P3/8/8/K7 w - - 0 1",
			to:       "7k/8/8/3P4/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "White pawn illegal backward capture",
			from:     "7k/8/8/8/4P3/3p4/8/K7 w - - 0 1",
			to:       "7k/8/8/8/8/3P4/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black pawn captures left e5-d4",
			from:     "7k/8/8/4p3/3P4/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/3p4/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black pawn captures right e5-f4",
			from:     "7k/8/8/4p3/5N2/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/5p2/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black pawn diagonal empty square",
			from:     "7k/8/8/4p3/8/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/3p4/8/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "Black pawn captures friendly piece",
			from:     "7k/8/8/4p3/3n4/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/3p4/8/8/K7 b - - 0 1",
			expected: false,
		}, {
			name:     "Vertical pinned white pawn single push",
			from:     "4r2k/8/8/8/8/8/4P3/4K3 w - - 0 1",
			to:       "4r2k/8/8/8/8/4P3/8/4K3 w - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pinned white pawn double push",
			from:     "4r2k/8/8/8/8/8/4P3/4K3 w - - 0 1",
			to:       "4r2k/8/8/8/4P3/8/8/4K3 w - e3 0 1",
			expected: true,
		},
		{
			name:     "Vertical pinned white pawn captures left",
			from:     "4r2k/8/8/8/8/3n4/4P3/4K3 w - - 0 1",
			to:       "4r2k/8/8/8/8/3P4/8/4K3 w - - 0 1",
			expected: false,
		},
		{
			name:     "Vertical pinned white pawn captures right",
			from:     "4r2k/8/8/8/8/5n2/4P3/4K3 w - - 0 1",
			to:       "4r2k/8/8/8/8/5P2/8/4K3 w - - 0 1",
			expected: false,
		},
		{
			name:     "Vertical pinned black pawn single push",
			from:     "4k3/4p3/8/8/8/8/8/K3R3 b - - 0 1",
			to:       "4k3/8/4p3/8/8/8/8/K3R3 b - - 0 1",
			expected: true,
		},
		{
			name:     "Vertical pinned black pawn double push",
			from:     "4k3/4p3/8/8/8/8/8/K3R3 b - - 0 1",
			to:       "4k3/8/8/4p3/8/8/8/K3R3 b - e6 0 1",
			expected: true,
		}, {
			name:     "Horizontal pinned white pawn forward",
			from:     "7k/8/8/8/r2P3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/3P4/r6K/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "Horizontal pinned white pawn captures",
			from:     "7k/8/8/4n3/r2P3K/8/8/8 w - - 0 1",
			to:       "7k/8/8/4P3/r6K/8/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "Horizontal pinned black pawn forward",
			from:     "8/8/8/k2p3R/8/8/8/7K b - - 0 1",
			to:       "8/8/8/k6R/3p4/8/8/7K b - - 0 1",
			expected: false,
		}, {
			name:     "NW pinned white pawn illegal forward",
			from:     "7k/8/8/8/8/3b4/4P3/5K2 w - - 0 1",
			to:       "7k/8/8/8/8/3bP3/8/5K2 w - - 0 1",
			expected: false,
		},
		{
			name:     "NW pinned white pawn captures pinning bishop",
			from:     "7k/8/8/8/8/3b4/4P3/5K2 w - - 0 1",
			to:       "7k/8/8/8/8/3P4/8/5K2 w - - 0 1",
			expected: true,
		},
		{
			name:     "NE pinned white pawn captures pinning bishop",
			from:     "7k/8/8/8/8/5b2/4P3/3K4 w - - 0 1",
			to:       "7k/8/8/8/8/5P2/8/3K4 w - - 0 1",
			expected: true,
		},
		{
			name:     "NE pinned white pawn captures outside pin",
			from:     "7k/8/8/8/8/3n1b2/4P3/3K4 w - - 0 1",
			to:       "7k/8/8/8/8/3P1b2/8/3K4 w - - 0 1",
			expected: false,
		},
		{
			name:     "SW pinned black pawn captures pinning bishop",
			from:     "5k2/4p3/3B4/8/8/8/8/K7 b - - 0 1",
			to:       "5k2/8/3p4/8/8/8/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "SE pinned black pawn captures pinning bishop",
			from:     "3k4/4p3/5B2/8/8/8/8/K7 b - - 0 1",
			to:       "3k4/8/5p2/8/8/8/8/K7 b - - 0 1",
			expected: true,
		}, {
			name:     "White en passant left e5-d6",
			from:     "7k/8/8/3pP3/8/8/8/K7 w - d6 0 1",
			to:       "7k/8/3P4/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "White en passant right e5-f6",
			from:     "7k/8/8/4Pp2/8/8/8/K7 w - f6 0 1",
			to:       "7k/8/5P2/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Black en passant left e4-d3",
			from:     "7k/8/8/8/3Pp3/8/8/K7 b - d3 0 1",
			to:       "7k/8/8/8/8/3p4/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black en passant right e4-f3",
			from:     "7k/8/8/8/4pP2/8/8/K7 b - f3 0 1",
			to:       "7k/8/8/8/8/5p2/8/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "White en passant edge file a5-b6",
			from:     "7k/8/8/Pp6/8/8/8/K7 w - b6 0 1",
			to:       "7k/8/1P6/8/8/8/8/K7 w - - 0 1",
			expected: true,
		},
		{
			name:     "Black en passant edge file h4-g3",
			from:     "7k/8/8/8/6Pp/8/8/K7 b - g3 0 1",
			to:       "7k/8/8/8/8/6p1/8/K7 b - - 0 1",
			expected: true,
		}, {
			name:     "White en passant without target",
			from:     "7k/8/8/3pP3/8/8/8/K7 w - - 0 1",
			to:       "7k/8/3P4/8/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black en passant without target",
			from:     "7k/8/8/8/3Pp3/8/8/K7 b - - 0 1",
			to:       "7k/8/8/8/8/3p4/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "White en passant wrong target",
			from:     "7k/8/8/3pPp2/8/8/8/K7 w - f6 0 1",
			to:       "7k/8/3P4/5p2/8/8/8/K7 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black en passant wrong target",
			from:     "7k/8/8/8/3PpP2/8/8/K7 b - f3 0 1",
			to:       "7k/8/8/8/5P2/3p4/8/K7 b - - 0 1",
			expected: false,
		},
		{
			name:     "White en passant from wrong rank",
			from:     "7k/8/8/3p4/4P3/8/8/K7 w - d6 0 1",
			to:       "7k/8/3P4/8/8/8/8/K7 w - - 0 1",
			expected: true,
		}, {
			name:     "White vertically pinned pawn en passant",
			from:     "4r2k/8/8/3pP3/8/8/8/4K3 w - d6 0 1",
			to:       "4r2k/8/3P4/8/8/8/8/4K3 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black vertically pinned pawn en passant",
			from:     "4k3/8/8/8/3Pp3/8/8/K3R3 b - d3 0 1",
			to:       "4k3/8/8/8/8/3p4/8/K3R3 b - - 0 1",
			expected: false,
		},

		{
			name:     "White en passant horizontal attack still blocked",
			from:     "7k/8/8/r1NpP2K/8/8/8/8 w - d6 0 1",
			to:       "7k/8/3P4/r1N4K/8/8/8/8 w - - 0 1",
			expected: true,
		}, {
			name:     "White diagonal pinned pawn en passant along pin",
			from:     "7k/2b5/8/3pP3/8/6K1/8/8 w - d6 0 1",
			to:       "7k/2b5/3P4/8/8/6K1/8/8 w - - 0 1",
			expected: true,
		},
		{
			name:     "White diagonal pinned pawn en passant breaks pin",
			from:     "7k/6b1/8/3pP3/8/2K5/8/8 w - d6 0 1",
			to:       "7k/6b1/3P4/8/8/2K5/8/8 w - - 0 1",
			expected: false,
		},
		{
			name:     "Black diagonal pinned pawn en passant along pin",
			from:     "8/8/6k1/8/3Pp3/8/2B5/K7 b - d3 0 1",
			to:       "8/8/6k1/8/8/3p4/2B5/K7 b - - 0 1",
			expected: true,
		},
		{
			name:     "Black diagonal pinned pawn en passant breaks pin",
			from:     "8/8/2k5/8/3Pp3/8/6B1/K7 b - d3 0 1",
			to:       "8/8/2k5/8/8/3p4/6B1/K7 b - - 0 1",
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

		valid, err := game.ValidatePawnMove(parseFrom, parseTo)
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
