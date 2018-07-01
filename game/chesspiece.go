// chesspiece
package chessgame

type Coordinate struct {
	Row    int
	Column int
}

type ChessPiece interface {
	validMoves(board *ChessBoard) map[Coordinate]bool
	updatePosition(coord Coordinate)
	getPieceSide() Side
	getPieceType() PieceType
	hasPieceMoved() bool
	getCurrentCoordinates() Coordinate
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
