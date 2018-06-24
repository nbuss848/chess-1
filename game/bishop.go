// bishop
package chessgame

type Bishop struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

func newBishop(bishopSide Side, coord Coordinate) Bishop {
	return Bishop{currentCoordinate: coord, pieceSide: bishopSide, hasMoved: false}
}

func (bishop *Bishop) updatePosition(coord Coordinate) {
	bishop.currentCoordinate = coord
	bishop.hasMoved = true
}

func (bishop *Bishop) validMoves(board *ChessBoard) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	allMoves := getAllDiagonalMoves(bishop.currentCoordinate, board, bishop.pieceSide)
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
