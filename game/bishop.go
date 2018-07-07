// bishop
package chessgame

type Bishop struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

// Sets hasMoved property in Bishop struct. Used for cloning
func (bishop *Bishop) setHasMoved(hasMoved bool) {
	bishop.hasMoved = hasMoved
}

func (bishop *Bishop) getCurrentCoordinates() Coordinate {
	return bishop.currentCoordinate
}

func newBishop(bishopSide Side, coord Coordinate) Bishop {
	return Bishop{currentCoordinate: coord, pieceSide: bishopSide, hasMoved: false}
}

func (bishop *Bishop) updatePosition(coord Coordinate) {
	bishop.currentCoordinate = coord
	bishop.hasMoved = true
}

func (bishop *Bishop) validMoves(board *ChessBoard) map[Coordinate]bool {
	return getAllMovesForPiece(board, bishop, getAllBishopMoves)
}

func getAllBishopMoves(board *ChessBoard, bishop ChessPiece) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	allMoves := getAllDiagonalMoves(bishop.getCurrentCoordinates(), board, bishop.getPieceSide())
	for i := 0; i < len(allMoves); i++ {
		validMoves[allMoves[i]] = true
	}
	return validMoves
}

func (bishop *Bishop) getPieceSide() Side {
	return bishop.pieceSide
}

func (bishop *Bishop) getPieceType() PieceType {
	return BISHOP
}

func (bishop *Bishop) hasPieceMoved() bool {
	return bishop.hasMoved
}
