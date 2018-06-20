// rook
package chessgame

type Rook struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
}

func (rook *Rook) updatePosition(coord Coordinate) {
	rook.currentCoordinate = coord
}

func (rook *Rook) updateValidMoves(board Board) {
	rook.potentialMoves = make(map[Coordinate]bool)
	allMoves := getAllStraightLineMoves(rook.currentCoordinate, board, rook.pieceSide)
	for i := 0; i < len(allMoves); i++ {
		rook.potentialMoves[allMoves[i]] = true
	}
}

func (rook *Rook) getPieceSide() {
	return rook.pieceSide
}

func (rook *Rook) validMoves() []Coordinate {
	var potentialMoves []Coordinate
	for k := range rook.potentialMoves {
		potentialMoves = append(potentialMoves, k)
	}
	return potentialMoves
}
