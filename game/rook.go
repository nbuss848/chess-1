// rook
package chessgame

type Rook struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	hasMoved          bool
}

// Sets hasMoved property in Rook struct. Used for cloning
func (rook *Rook) setHasMoved(hasMoved bool) {
	rook.hasMoved = hasMoved
}

func (rook *Rook) GetCurrentCoordinates() Coordinate {
	return rook.currentCoordinate
}

func newRook(rookSide Side, coord Coordinate) Rook {
	return Rook{currentCoordinate: coord, pieceSide: rookSide, hasMoved: false}
}

func (rook *Rook) updatePosition(coord Coordinate) {
	rook.currentCoordinate = coord
	rook.hasMoved = true
}

func (rook *Rook) ValidMoves(board *ChessBoard) map[Coordinate]bool {
	return getAllMovesForPiece(board, rook, getAllRookMoves)
}

func getAllRookMoves(board *ChessBoard, rook ChessPiece) map[Coordinate]bool {
	allMovesSlice := getAllStraightLineMoves(rook.GetCurrentCoordinates(), board, rook.GetPieceSide())
	moveMap := make(map[Coordinate]bool)
	for i := 0; i < len(allMovesSlice); i++ {
		moveMap[allMovesSlice[i]] = true
	}
	return moveMap
}

func (rook *Rook) GetPieceSide() Side {
	return rook.pieceSide
}

func (rook *Rook) GetPieceType() PieceType {
	return ROOK
}

func (rook *Rook) hasPieceMoved() bool {
	return rook.hasMoved
}
