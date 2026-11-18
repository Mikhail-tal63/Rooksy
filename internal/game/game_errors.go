package game

import "net/http"

type GameErrors struct {
	msg    string
	status int
	code   string
}

func (g *GameErrors) Error() string   { return g.msg }
func (g *GameErrors) Code() string    { return g.code }
func (g *GameErrors) HTTPStatus() int { return g.status }

var (
	// FEN errors
	ErrInvalidFEN             error = &GameErrors{"invalid FEN notation", http.StatusBadRequest, "INVALID_FEN"}
	ErrInvalidFENBoard        error = &GameErrors{"invalid FEN board representation", http.StatusBadRequest, "INVALID_FEN_BOARD"}
	ErrInvalidActiveColor     error = &GameErrors{"invalid active color in FEN", http.StatusBadRequest, "INVALID_ACTIVE_COLOR"}
	ErrInvalidCastlingRights  error = &GameErrors{"invalid castling rights in FEN", http.StatusBadRequest, "INVALID_CASTLING_RIGHTS"}
	ErrInvalidEnPassantSquare error = &GameErrors{"invalid en passant square in FEN", http.StatusBadRequest, "INVALID_EN_PASSANT_SQUARE"}
	ErrInvalidHalfmoveClock   error = &GameErrors{"invalid halfmove clock in FEN", http.StatusBadRequest, "INVALID_HALFMOVE_CLOCK"}
	ErrInvalidFullmoveNumber  error = &GameErrors{"invalid fullmove number in FEN", http.StatusBadRequest, "INVALID_FULLMOVE_NUMBER"}

	// Board and piece errors
	ErrInvalidPiece      error = &GameErrors{"invalid piece", http.StatusBadRequest, "INVALID_PIECE"}
	ErrInvalidPieceColor error = &GameErrors{"invalid piece color", http.StatusBadRequest, "INVALID_PIECE_COLOR"}
	ErrInvalidPieceType  error = &GameErrors{"invalid piece type", http.StatusBadRequest, "INVALID_PIECE_TYPE"}
	ErrInvalidSquare     error = &GameErrors{"invalid square", http.StatusBadRequest, "INVALID_SQUARE"}
	ErrInvalidBoard      error = &GameErrors{"invalid chessboard state", http.StatusBadRequest, "INVALID_BOARD"}
	ErrInvalidKingCount  error = &GameErrors{"invalid number of kings", http.StatusBadRequest, "INVALID_KING_COUNT"}

	// Move errors
	ErrInvalidMove           error = &GameErrors{"invalid move", http.StatusBadRequest, "INVALID_MOVE"}
	ErrIllegalMove           error = &GameErrors{"illegal move", http.StatusBadRequest, "ILLEGAL_MOVE"}
	ErrWrongTurn             error = &GameErrors{"not your turn", http.StatusForbidden, "WRONG_TURN"}
	ErrOwnPieceAtDestination error = &GameErrors{"destination square contains your own piece", http.StatusBadRequest, "OWN_PIECE_AT_DESTINATION"}
	ErrKingInCheck           error = &GameErrors{"move leaves your king in check", http.StatusBadRequest, "KING_IN_CHECK"}
	ErrInvalidPromotion      error = &GameErrors{"invalid pawn promotion", http.StatusBadRequest, "INVALID_PROMOTION"}
	ErrInvalidCastling       error = &GameErrors{"invalid castling move", http.StatusBadRequest, "INVALID_CASTLING"}
	ErrInvalidEnPassant      error = &GameErrors{"invalid en passant move", http.StatusBadRequest, "INVALID_EN_PASSANT"}

	// Game lifecycle errors
	ErrGameNotFound        error = &GameErrors{"game not found", http.StatusNotFound, "GAME_NOT_FOUND"}
	ErrGameAlreadyStarted  error = &GameErrors{"game has already started", http.StatusConflict, "GAME_ALREADY_STARTED"}
	ErrGameNotStarted      error = &GameErrors{"game has not started", http.StatusConflict, "GAME_NOT_STARTED"}
	ErrGameAlreadyFinished error = &GameErrors{"game has already finished", http.StatusConflict, "GAME_ALREADY_FINISHED"}
	ErrGameAbandoned       error = &GameErrors{"game has been abandoned", http.StatusConflict, "GAME_ABANDONED"}
	ErrGameNotActive       error = &GameErrors{"game is not active", http.StatusConflict, "GAME_NOT_ACTIVE"}
	ErrPlayerNotInGame     error = &GameErrors{"player is not part of this game", http.StatusForbidden, "PLAYER_NOT_IN_GAME"}
	ErrPlayerAlreadyInGame error = &GameErrors{"player is already part of this game", http.StatusConflict, "PLAYER_ALREADY_IN_GAME"}
	ErrSamePlayer          error = &GameErrors{"a player cannot play against themselves", http.StatusBadRequest, "SAME_PLAYER"}

	// Game result errors
	ErrInvalidGameResult   error = &GameErrors{"invalid game result", http.StatusBadRequest, "INVALID_GAME_RESULT"}
	ErrGameAlreadyResigned error = &GameErrors{"player has already resigned", http.StatusConflict, "GAME_ALREADY_RESIGNED"}
	ErrDrawNotAvailable    error = &GameErrors{"draw is not available", http.StatusConflict, "DRAW_NOT_AVAILABLE"}
	ErrDrawAlreadyOffered  error = &GameErrors{"a draw has already been offered", http.StatusConflict, "DRAW_ALREADY_OFFERED"}

	// Clock errors
	ErrInvalidTimeControl  error = &GameErrors{"invalid time control", http.StatusBadRequest, "INVALID_TIME_CONTROL"}
	ErrClockNotRunning     error = &GameErrors{"game clock is not running", http.StatusConflict, "CLOCK_NOT_RUNNING"}
	ErrClockAlreadyRunning error = &GameErrors{"game clock is already running", http.StatusConflict, "CLOCK_ALREADY_RUNNING"}
	ErrTimeExpired         error = &GameErrors{"player's time has expired", http.StatusConflict, "TIME_EXPIRED"}

	// History and state errors
	ErrInvalidMoveHistory   error = &GameErrors{"invalid move history", http.StatusBadRequest, "INVALID_MOVE_HISTORY"}
	ErrGameStateMismatch    error = &GameErrors{"game state does not match move history", http.StatusConflict, "GAME_STATE_MISMATCH"}
	ErrConcurrentGameUpdate error = &GameErrors{"game was updated by another operation", http.StatusConflict, "CONCURRENT_GAME_UPDATE"}
)
