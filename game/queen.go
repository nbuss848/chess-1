// queen
package chessgame

type Queen struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	hasMoved          bool
}

func newQueen(queenSide Side, coord Coordinate) Queen {
	return Queen{currentCoordinate: coord, pieceSide: queenSide, hasMoved: false}
}

func (queen *Queen) updatePosition(coord Coordinate) {
	queen.currentCoordinate = coord
	queen.hasMoved = true
}

func (queen *Queen) updateValidMoves(board *ChessBoard) {
	queen.potentialMoves = make(map[Coordinate]bool)
	allStraightLineMoves := getAllStraightLineMoves(queen.currentCoordinate, board, queen.pieceSide)
	for i := 0; i < len(allStraightLineMoves); i++ {
		queen.potentialMoves[allStraightLineMoves[i]] = true
	}
	allDiagonalMoves := getAllDiagonalMoves(queen.currentCoordinate, board, queen.pieceSide)
	for i := 0; i < len(allDiagonalMoves); i++ {
		queen.potentialMoves[allDiagonalMoves[i]] = true
	}
}

func (queen *Queen) getPieceSide() Side {
	return queen.pieceSide
}

func (queen *Queen) validMoves() []Coordinate {
	var potentialMoves []Coordinate
	for k := range queen.potentialMoves {
		potentialMoves = append(potentialMoves, k)
	}
	return potentialMoves
}

func (queen *Queen) getPieceType() PieceType {
	return QUEEN
}

func (queen *Queen) hasPieceMoved() bool {
	return queen.hasMoved
}
