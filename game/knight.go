// knight
package chessgame

type Knight struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	hasMoved          bool
}

// Sets hasMoved property in Knight struct. Used for cloning
func (knight *Knight) setHasMoved(hasMoved bool) {
	knight.hasMoved = hasMoved
}

func (knight *Knight) GetCurrentCoordinates() Coordinate {
	return knight.currentCoordinate
}

func newKnight(knightSide Side, coord Coordinate) Knight {
	return Knight{currentCoordinate: coord, pieceSide: knightSide, hasMoved: false}
}

func (knight *Knight) updatePosition(coord Coordinate) {
	knight.currentCoordinate = coord
	knight.hasMoved = true
}

func (knight *Knight) ValidMoves(board *ChessBoard) map[Coordinate]bool {
	return getAllMovesForPiece(board, knight, getAllKnightMoves)
}

func getAllKnightMoves(board *ChessBoard, knight ChessPiece) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	allPotentialCoordinates := getAllPossibleKnightMoves(knight.GetCurrentCoordinates())
	for i := 0; i < len(allPotentialCoordinates); i++ {
		canMove, _ := canMoveToSquare(allPotentialCoordinates[i], board, knight.GetPieceSide())
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

func (knight *Knight) GetPieceSide() Side {
	return knight.pieceSide
}

func (knight *Knight) GetPieceType() PieceType {
	return KNIGHT
}

func (knight *Knight) hasPieceMoved() bool {
	return knight.hasMoved
}
