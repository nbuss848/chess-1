// queen
package chessgame

type Queen struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

// Sets hasMoved property in Queen struct. Used for cloning
func (queen *Queen) setHasMoved(hasMoved bool) {
	queen.hasMoved = hasMoved
}

func (queen *Queen) GetCurrentCoordinates() Coordinate {
	return queen.currentCoordinate
}

func newQueen(queenSide Side, coord Coordinate) Queen {
	return Queen{currentCoordinate: coord, pieceSide: queenSide, hasMoved: false}
}

func (queen *Queen) updatePosition(coord Coordinate) {
	queen.currentCoordinate = coord
	queen.hasMoved = true
}

func (queen *Queen) ValidMoves(board *ChessBoard) map[Coordinate]bool {
	return getAllMovesForPiece(board, queen, getAllQueenMoves)
}

func getAllQueenMoves(board *ChessBoard, queen ChessPiece) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	allStraightLineMoves := getAllStraightLineMoves(queen.GetCurrentCoordinates(), board, queen.GetPieceSide())
	for i := 0; i < len(allStraightLineMoves); i++ {
		validMoves[allStraightLineMoves[i]] = true
	}
	allDiagonalMoves := getAllDiagonalMoves(queen.GetCurrentCoordinates(), board, queen.GetPieceSide())
	for i := 0; i < len(allDiagonalMoves); i++ {
		validMoves[allDiagonalMoves[i]] = true
	}
	return validMoves
}

func (queen *Queen) GetPieceSide() Side {
	return queen.pieceSide
}

func (queen *Queen) GetPieceType() PieceType {
	return QUEEN
}

func (queen *Queen) hasPieceMoved() bool {
	return queen.hasMoved
}
