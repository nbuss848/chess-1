// chesspiece
package chessgame

type Coordinate struct {
	Row    int
	Column int
}

type ChessPiece interface {
	validMoves(board Board) []Coordinate
	updatePosition(coordinate Coordinate)
	getPieceSide(coord Coordinate) Side
}

type Side int

const (
	WHITE Side = 1 + iota
	BLACK
)
