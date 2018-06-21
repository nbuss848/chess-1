// chessboard
package chessgame

type BoardData struct {
	BoardPieces [][]ChessPiece
	MoveLog     []Move
}

type Move struct {
	fromCoordinate Coordinate
	toCoordinate   Coordinate
	isCapture      bool
	piece          ChessPiece
}

type Board interface {
	isSpaceOccupied(coord Coordinate) bool
	getPieceSide(coord Coordinate) Side
	getPreviousMove() Move
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
	return board.BoardPieces[coord.Row][coord.Column].getPieceSide(coord)
}

func (board BoardData) getPreviousMove() Move {
	return board.MoveLog[len(board.MoveLog)-1]
}
