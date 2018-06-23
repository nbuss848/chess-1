// rook
package chessgame

type Rook struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	hasMoved          bool
}

func newRook(rookSide Side, coord Coordinate) Rook {
	return Rook{currentCoordinate: coord, pieceSide: rookSide, hasMoved: false}
}

func (rook *Rook) updatePosition(coord Coordinate) {
	rook.currentCoordinate = coord
	rook.hasMoved = true
}

func (rook *Rook) updateValidMoves(board *ChessBoard) {
	rook.potentialMoves = make(map[Coordinate]bool)
	allMoves := getAllStraightLineMoves(rook.currentCoordinate, board, rook.pieceSide)
	for i := 0; i < len(allMoves); i++ {
		rook.potentialMoves[allMoves[i]] = true
	}
}

func (rook *Rook) getPieceSide() Side {
	return rook.pieceSide
}

func (rook *Rook) validMoves() []Coordinate {
	var potentialMoves []Coordinate
	for k := range rook.potentialMoves {
		potentialMoves = append(potentialMoves, k)
	}
	return potentialMoves
}

func (rook *Rook) getPieceType() PieceType {
	return ROOK
}

func (rook *Rook) hasPieceMoved() bool {
	return rook.hasMoved
}
