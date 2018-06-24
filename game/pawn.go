// pawn
package chessgame

type Pawn struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

func newPawn(pawnSide Side, coord Coordinate) Pawn {
	return Pawn{currentCoordinate: coord, pieceSide: pawnSide, hasMoved: false}
}

func (pawn *Pawn) updatePosition(coord Coordinate) {
	pawn.currentCoordinate = coord
	pawn.hasMoved = true
}

func (pawn *Pawn) validMoves(board *ChessBoard) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	moveChange := 1
	if pawn.pieceSide == BLACK {
		moveChange = -1
	}
	oneMovePotentialCoordinate := Coordinate{Row: pawn.currentCoordinate.Row + moveChange, Column: pawn.currentCoordinate.Column}
	if !board.isSpaceOccupied(oneMovePotentialCoordinate) {
		validMoves[oneMovePotentialCoordinate] = true
	}
	twoMovePotentialCoordinate := Coordinate{Row: oneMovePotentialCoordinate.Row + moveChange, Column: oneMovePotentialCoordinate.Column}
	if !pawn.hasMoved && !board.isSpaceOccupied((twoMovePotentialCoordinate)) {
		validMoves[twoMovePotentialCoordinate] = true
	}
	firstCaptureMove := Coordinate{Row: oneMovePotentialCoordinate.Row, Column: oneMovePotentialCoordinate.Column + 1}
	if validateCaptureMove(board, firstCaptureMove, pawn.pieceSide, false) {
		validMoves[firstCaptureMove] = true
	}
	secondCaptureMove := Coordinate{Row: oneMovePotentialCoordinate.Row, Column: oneMovePotentialCoordinate.Column - 1}
	if validateCaptureMove(board, firstCaptureMove, pawn.pieceSide, true) {
		validMoves[secondCaptureMove] = true
	}
	lastMove, wasLastMove := board.getPreviousMove()
	if wasLastMove == false {
		return validMoves
	}
	if lastMove.piece.getPieceType() == PAWN && AbsIntVal(lastMove.fromCoordinate.Row-lastMove.toCoordinate.Row) == 2 && lastMove.toCoordinate.Row == pawn.currentCoordinate.Row {
		if lastMove.fromCoordinate.Column-pawn.currentCoordinate.Column == -1 {
			lowerColEnPassant := Coordinate{Row: pawn.currentCoordinate.Row + moveChange, Column: pawn.currentCoordinate.Column - 1}
			validMoves[lowerColEnPassant] = true
		} else if lastMove.fromCoordinate.Column-pawn.currentCoordinate.Column == 1 {
			higherColEnPassant := Coordinate{Row: pawn.currentCoordinate.Row + moveChange, Column: pawn.currentCoordinate.Column + 1}
			validMoves[higherColEnPassant] = true
		}
	}
	return validMoves
}

func (pawn *Pawn) getPieceSide() Side {
	return pawn.pieceSide
}

func (pawn *Pawn) getPieceType() PieceType {
	return PAWN
}

func (pawn *Pawn) hasPieceMoved() bool {
	return pawn.hasMoved
}

func validateCaptureMove(board *ChessBoard, coord Coordinate, pieceSide Side, left bool) bool {
	if !coord.isLegal() {
		return false
	}
	if !board.isSpaceOccupied(coord) {
		return false
	}
	if board.getPieceSide(coord) != pieceSide {
		return false
	}
	return true
}
