// chessboard
package chessgame

type ChessBoard struct {
	BoardPieces [][]ChessPiece
	MoveLog     []Move
}

type Move struct {
	fromCoordinate Coordinate
	toCoordinate   Coordinate
	isCapture      bool
	piece          ChessPiece
}

// TODO: combine isSpaceOccupied with getPieceSide
func (board *ChessBoard) isSpaceOccupied(coord Coordinate) bool {
	if board.BoardPieces[coord.Row][coord.Column] != nil {
		return true
	}
	return false
}

func (board *ChessBoard) getPieceType(coord Coordinate) PieceType {
	return board.getPieceType(coord)
}

func (board *ChessBoard) getPieceSide(coord Coordinate) Side {
	return board.BoardPieces[coord.Row][coord.Column].getPieceSide(coord)
}

func (board *ChessBoard) getPreviousMove() Move {
	return board.MoveLog[len(board.MoveLog)-1]
}

func (board *ChessBoard) hasPieceMoved(coord Coordinate) bool {
	return board.BoardPieces[coord.Row][coord.Column].hasPieceMoved()
}
