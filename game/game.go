// game
package chessgame

// Gives outcome of a move
type TurnOutcome int

const (
	WHITECHECKMATE TurnOutcome = 1 + iota
	BLACKCHECKMATE
	STALEMATE
	CONTINUE
)

// Gives outcome of game
type GameOutcome int

const (
	WHITEVICTORY GameOutcome = 1 + iota
	BLACKVICTORY
	DRAW
)

// Chess player interface - given board, takes turn, and returns outcome of that turn.
type ChessPlayer interface {
	MakeMove(boardClone ChessBoard, validMoves map[Coordinate]map[Coordinate]bool) (Coordinate, Coordinate)
	PromotePawn() PieceType
	GetSide() Side
}

// Chess game, has player who plays on white side (moves first) and player who plays on black side (moves second)
// pointer to board, which is used to represent internal state of game
type ChessGame struct {
	WhitePlayer ChessPlayer
	BlackPlayer ChessPlayer
	Board       *ChessBoard
}

// Given two players, returns new chess game
func NewChessGame(whitePlayer ChessPlayer, blackPlayer ChessPlayer) ChessGame {
	board := CreateBoard()
	return ChessGame{whitePlayer, blackPlayer, &board}
}

// Loops infinitely, having players take turn at each pass, until outcome returned isn't CONTINUE
func (game ChessGame) PlayGame() GameOutcome {
	var whiteTurn TurnOutcome
	var blackTurn TurnOutcome
	for {
		whiteTurn = game.Board.TakeTurn(game.WhitePlayer)
		if whiteTurn != CONTINUE {
			break
		}
		blackTurn = game.Board.TakeTurn(game.BlackPlayer)
		if blackTurn != CONTINUE {
			break
		}
	}
	if whiteTurn != CONTINUE {
		return turnOutcomeToGameOutcome(whiteTurn)
	} else {
		return turnOutcomeToGameOutcome(blackTurn)
	}
}

// Converts turn outcome into game outcome - turn outcome should not be continue
func turnOutcomeToGameOutcome(turnOutcome TurnOutcome) GameOutcome {
	if turnOutcome == BLACKCHECKMATE {
		return BLACKVICTORY
	} else if turnOutcome == WHITECHECKMATE {
		return WHITEVICTORY
	} else {
		return DRAW
	}
}
