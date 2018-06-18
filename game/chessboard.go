// chessboard
package chessgame

type BoardData struct {
	BoardPieces [][]ChessPiece
}

type Board interface {
	isSpaceOccupied(coord Coordinate) bool
	getPieceSide(coord Coordinate) Side
}

// TODO: combine isSpaceOccupied with getPieceSide
func (board BoardData) isSpaceOccupied(coord Coordinate) bool {
	if board.BoardPieces[coord.Row][coord.Column] != nil {
		return true
	}
	return false
}

// TODO: research how enumerators work in Go
func (board BoardData) getPieceSide(coord Coordinate) Side {
	return board.getPieceSide(coord)
}
