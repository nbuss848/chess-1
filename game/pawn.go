// pawn
package chessgame

type Pawn struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	hasMoved          bool
}

func (pawn *Pawn) updatePosition(coord Coordinate) {
	pawn.currentCoordinate = coord
	pawn.hasMoved = true
}

func (pawn *Pawn) updateValidMoves(board *ChessBoard) {
	pawn.potentialMoves = make(map[Coordinate]bool)
	moveChange := 1
	if pawn.pieceSide == BLACK {
		moveChange = -1
	}
	oneMovePotentialCoordinate := Coordinate{Row: pawn.currentCoordinate.Row + moveChange, Column: pawn.currentCoordinate.Column}
	if !board.isSpaceOccupied(oneMovePotentialCoordinate) {
		pawn.potentialMoves[oneMovePotentialCoordinate] = true
	}
	twoMovePotentialCoordinate := Coordinate{Row: oneMovePotentialCoordinate.Row + moveChange, Column: oneMovePotentialCoordinate.Column}
	if !pawn.hasMoved && !board.isSpaceOccupied((twoMovePotentialCoordinate)) {
		pawn.potentialMoves[twoMovePotentialCoordinate] = true
	}
	firstCaptureMove := Coordinate{Row: oneMovePotentialCoordinate.Row, Column: oneMovePotentialCoordinate.Column + 1}
	if board.isSpaceOccupied(firstCaptureMove) && board.getPieceSide(firstCaptureMove) != pawn.getPieceSide() {
		pawn.potentialMoves[firstCaptureMove] = true
	}
	secondCaptureMove := Coordinate{Row: oneMovePotentialCoordinate.Row, Column: oneMovePotentialCoordinate.Column - 1}
	if board.isSpaceOccupied(secondCaptureMove) && board.getPieceSide(secondCaptureMove) != pawn.getPieceSide() {
		pawn.potentialMoves[secondCaptureMove] = true
	}
	lastMove := board.getPreviousMove()
	if lastMove.piece.getPieceType() == PAWN && AbsIntVal(lastMove.fromCoordinate.Row-lastMove.toCoordinate.Row) == 2 && lastMove.toCoordinate.Row == pawn.currentCoordinate.Row {
		if lastMove.fromCoordinate.Column-pawn.currentCoordinate.Column == -1 {
			lowerColEnPassant := Coordinate{Row: pawn.currentCoordinate.Row + moveChange, Column: pawn.currentCoordinate.Column - 1}
			pawn.potentialMoves[lowerColEnPassant] = true
		} else if lastMove.fromCoordinate.Column-pawn.currentCoordinate.Column == 1 {
			higherColEnPassant := Coordinate{Row: pawn.currentCoordinate.Row + moveChange, Column: pawn.currentCoordinate.Column + 1}
			pawn.potentialMoves[higherColEnPassant] = true
		}
	}
}

func (pawn *Pawn) getPieceSide() Side {
	return pawn.pieceSide
}

func (pawn *Pawn) validMoves() []Coordinate {
	var potentialMoves []Coordinate
	for k := range pawn.potentialMoves {
		potentialMoves = append(potentialMoves, k)
	}
	return potentialMoves
}

func (pawn *Pawn) getPieceType() PieceType {
	return PAWN
}

func (pawn *Pawn) hasPieceMoved() bool {
	return pawn.hasMoved
}
