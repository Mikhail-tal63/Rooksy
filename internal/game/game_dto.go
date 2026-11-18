package game

type FEN struct {
	Board [8][8]Piece

	ActiveColor   Colors
	Castling      CastlingRights
	EnPassant     string
	FiftyMove     int
	FullmoveCount int
}

type Colors int

const (
	White Colors = iota
	Black
)

type Square int

const (
	Empty Square = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Piece struct {
	Type  Square
	Color Colors
}

type CastlingRights struct {
	WhiteKingSide  bool
	WhiteQueenSide bool
	BlackKingSide  bool
	BlackQueenSide bool
}
