// bishop
package chessgame

type Bishop struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	hasMoved          bool
}

func (bishop *Bishop) updatePosition(coord Coordinate) {
	bishop.currentCoordinate = coord
	bishop.hasMoved = true
}

func (bishop *Bishop) updateValidMoves(board *ChessBoard) {
	bishop.potentialMoves = make(map[Coordinate]bool)
	allMoves := getAllDiagonalMoves(bishop.currentCoordinate, board, bishop.pieceSide)
	for i := 0; i < len(allMoves); i++ {
		bishop.potentialMoves[allMoves[i]] = true
	}
}

func (bishop *Bishop) getPieceSide() Side {
	return bishop.pieceSide
}

func (bishop *Bishop) validMoves() []Coordinate {
	var potentialMoves []Coordinate
	for k := range bishop.potentialMoves {
		potentialMoves = append(potentialMoves, k)
	}
	return potentialMoves
}

func (bishop *Bishop) getPieceType() PieceType {
	return BISHOP
}

func (bishop *Bishop) hasPieceMoved() bool {
	return bishop.hasMoved
}
