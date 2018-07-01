// knight
package chessgame

type Knight struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

func (knight *Knight) getCurrentCoordinates() Coordinate {
	return knight.currentCoordinate
}

func newKnight(knightSide Side, coord Coordinate) Knight {
	return Knight{currentCoordinate: coord, pieceSide: knightSide, hasMoved: false}
}

func (knight *Knight) updatePosition(coord Coordinate) {
	knight.currentCoordinate = coord
	knight.hasMoved = true
}

func (knight *Knight) validMoves(board *ChessBoard) map[Coordinate]bool {
	return getAllMovesForPiece(board, knight, getAllKnightMoves)
}

func getAllKnightMoves(board *ChessBoard, knight ChessPiece) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	allPotentialCoordinates := getAllPossibleKnightMoves(knight.getCurrentCoordinates())
	for i := 0; i < len(allPotentialCoordinates); i++ {
		canMove, _ := canMoveToSquare(allPotentialCoordinates[i], board, knight.getPieceSide())
		if canMove {
			validMoves[allPotentialCoordinates[i]] = true
		}
	}
	return validMoves
}

func getAllPossibleKnightMoves(coord Coordinate) []Coordinate {
	var coordinates []Coordinate
	for i := -2; i <= 2; i += 4 {
		newRow := coord.Row + i
		newColumn := coord.Column + i
		coordinates = append(coordinates, Coordinate{Row: newRow, Column: coord.Column - 1})
		coordinates = append(coordinates, Coordinate{Row: newRow, Column: coord.Column + 1})
		coordinates = append(coordinates, Coordinate{Row: coord.Row - 1, Column: newColumn})
		coordinates = append(coordinates, Coordinate{Row: coord.Row + 1, Column: newColumn})
	}
	return coordinates
}

func (knight *Knight) getPieceSide() Side {
	return knight.pieceSide
}

func (knight *Knight) getPieceType() PieceType {
	return KNIGHT
}

func (knight *Knight) hasPieceMoved() bool {
	return knight.hasMoved
}
