// pawn
package chessgame

type Pawn struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

// Sets hasMoved property in Pawn struct. Used for cloning
func (pawn *Pawn) setHasMoved(hasMoved bool) {
	pawn.hasMoved = hasMoved
}

func (pawn *Pawn) getCurrentCoordinates() Coordinate {
	return pawn.currentCoordinate
}

func newPawn(pawnSide Side, coord Coordinate) Pawn {
	return Pawn{currentCoordinate: coord, pieceSide: pawnSide, hasMoved: false}
}

func (pawn *Pawn) updatePosition(coord Coordinate) {
	pawn.currentCoordinate = coord
	pawn.hasMoved = true
}

func (pawn *Pawn) validMoves(board *ChessBoard) map[Coordinate]bool {
	return getAllMovesForPiece(board, pawn, getAllPawnMoves)
}

func getAllPawnMoves(board *ChessBoard, pawn ChessPiece) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	coords := pawn.getCurrentCoordinates()
	pieceSide := pawn.getPieceSide()
	moveChange := 1
	if pieceSide == BLACK {
		moveChange = -1
	}
	oneMovePotentialCoordinate := Coordinate{Row: coords.Row + moveChange, Column: coords.Column}
	if !board.isSpaceOccupied(oneMovePotentialCoordinate) {
		validMoves[oneMovePotentialCoordinate] = true
	}
	twoMovePotentialCoordinate := Coordinate{Row: oneMovePotentialCoordinate.Row + moveChange, Column: oneMovePotentialCoordinate.Column}
	if !pawn.hasPieceMoved() && !board.isSpaceOccupied((twoMovePotentialCoordinate)) {
		validMoves[twoMovePotentialCoordinate] = true
	}
	firstCaptureMove := Coordinate{Row: oneMovePotentialCoordinate.Row, Column: oneMovePotentialCoordinate.Column + 1}
	if validateCaptureMove(board, firstCaptureMove, pieceSide, false) {
		validMoves[firstCaptureMove] = true
	}
	secondCaptureMove := Coordinate{Row: oneMovePotentialCoordinate.Row, Column: oneMovePotentialCoordinate.Column - 1}
	if validateCaptureMove(board, secondCaptureMove, pieceSide, true) {
		validMoves[secondCaptureMove] = true
	}
	lastMove, wasLastMove := board.getPreviousMove()
	if wasLastMove == false {
		return validMoves
	}
	if lastMove.pieceType == PAWN && AbsIntVal(lastMove.fromCoordinate.Row-lastMove.toCoordinate.Row) == 2 && lastMove.toCoordinate.Row == coords.Row {
		if lastMove.fromCoordinate.Column-coords.Column == -1 {
			lowerColEnPassant := Coordinate{Row: coords.Row + moveChange, Column: coords.Column - 1}
			validMoves[lowerColEnPassant] = true
		} else if lastMove.fromCoordinate.Column-coords.Column == 1 {
			higherColEnPassant := Coordinate{Row: coords.Row + moveChange, Column: coords.Column + 1}
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
	if board.getPieceSide(coord) == pieceSide {
		return false
	}
	return true
}
