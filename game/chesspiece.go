// chesspiece
package chessgame

type Coordinate struct {
	Row    int
	Column int
}

type ChessPiece interface {
	validMoves() []Coordinate
	updatePosition(coord Coordinate)
	updateValidMoves(board *Board)
	getPieceSide(coord Coordinate) Side
}

type Side int

const (
	WHITE Side = 1 + iota
	BLACK
)
