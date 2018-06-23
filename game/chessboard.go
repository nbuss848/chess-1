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
	getPieceType(coord Coordinate) PieceType
}

// TODO: combine isSpaceOccupied with getPieceSide
func (board BoardData) isSpaceOccupied(coord Coordinate) bool {
	if board.BoardPieces[coord.Row][coord.Column] != nil {
		return true
	}
	return false
}

func (board BoardData) getPieceType(coord Coordinate) PieceType {
	return board.getPieceType(coord)
}

func (board BoardData) getPieceSide(coord Coordinate) Side {
	return board.BoardPieces[coord.Row][coord.Column].getPieceSide(coord)
}

func (board BoardData) getPreviousMove() Move {
	return board.MoveLog[len(board.MoveLog)-1]
}
