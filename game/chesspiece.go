// chesspiece
package chessgame

type Coordinate struct {
	Row    int
	Column int
}

type ChessPiece interface {
	validMoves() []Coordinate
	updatePosition(coord Coordinate)
	updateValidMoves(board *ChessBoard)
	getPieceSide(coord Coordinate) Side
	getPieceType() PieceType
	hasPieceMoved() bool
}

type Side int

const (
	WHITE Side = 1 + iota
	BLACK
)

type PieceType int

const (
	PAWN PieceType = 1 + iota
	BISHOP
	ROOK
	KNIGHT
	KING
	QUEEN
)
