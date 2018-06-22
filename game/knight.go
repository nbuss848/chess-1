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
	for i := -2; i <= 2; i += 4 {
		newRow := knight.currentCoordinate.Row + i
		newColumn := knight.currentCoordinate.Column + 1

		newVertCoordLeft := Coordinate{Row: newRow, Column: knight.currentCoordinate.Column - 1}
		newVertCoordRight := Coordinate{Row: newRow, Column: knight.currentCoordinate.Column + 1}
		newHorizontalCoordUp := Coordinate{Row: knight.currentCoordinate.Row - 1, Column: newColumn}
		newHorizontalCoordDown := Coordinate{Row: knight.currentCoordinate.Row + 1, Column: newColumn}

		canMoveLeft, _ := canMoveToSquare(newVertCoordLeft, board, knight.pieceSide)
		if canMoveLeft {
			knight.potentialMoves[newVertCoordLeft] = true
		}
		canMoveRight, _ := canMoveToSquare(newVertCoordRight, board, knight.pieceSide)
		if canMoveRight {
			knight.potentialMoves[newVertCoordRight] = true
		}
		canMoveUp, _ := canMoveToSquare(newHorizontalCoordUp, board, knight.pieceSide)
		if canMoveUp {
			knight.potentialMoves[newHorizontalCoordUp] = true
		}
		canMoveDown, _ := canMoveToSquare(newHorizontalCoordDown, board, knight.pieceSide)
		if canMoveDown {
			knight.potentialMoves[newHorizontalCoordDown] = true
		}
	}
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
