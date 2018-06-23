// knight
package chessgame

type Knight struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
}

func (knight *Knight) updatePosition(coord Coordinate) {
	knight.currentCoordinate = coord
}

func (knight *Knight) updateValidMoves(board Board) {
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
