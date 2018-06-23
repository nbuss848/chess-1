// knight
package chessgame

type Knight struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	hasMoved          bool
}

func newKnight(knightSide Side, coord Coordinate) Knight {
	return Knight{currentCoordinate: coord, pieceSide: knightSide, hasMoved: false}
}

func (knight *Knight) updatePosition(coord Coordinate) {
	knight.currentCoordinate = coord
	knight.hasMoved = true
}

func (knight *Knight) updateValidMoves(board *ChessBoard) {
	knight.potentialMoves = make(map[Coordinate]bool)
	allPotentialCoordinates := getAllPossibleKnightMoves(knight.currentCoordinate)
	for i := 0; i < len(allPotentialCoordinates); i++ {
		canMove, _ := canMoveToSquare(allPotentialCoordinates[i], board, knight.pieceSide)
		if canMove {
			knight.potentialMoves[allPotentialCoordinates[i]] = true
		}
	}
}

func getAllPossibleKnightMoves(coord Coordinate) []Coordinate {
	var coordinates []Coordinate
	for i := -2; i <= 2; i += 4 {
		newRow := coord.Row + i
		newColumn := coord.Column + 1
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

func (knight *Knight) validMoves() []Coordinate {
	var potentialMoves []Coordinate
	for k := range knight.potentialMoves {
		potentialMoves = append(potentialMoves, k)
	}
	return potentialMoves
}

func (knight *Knight) getPieceType() PieceType {
	return KNIGHT
}

func (knight *Knight) hasPieceMoved() bool {
	return knight.hasMoved
}
