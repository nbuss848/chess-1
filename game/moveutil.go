// moveutil
package chessgame

// Gets all straight line moves, given a pieces coordinates, its side, and a board. Used for rooks and queens
func getAllStraightLineMoves(coord Coordinate, board Board, side Side) []Coordinate {
	var potentialMoves []Coordinate
	potentialUpMoves := getStraightLineMoves(coord, board, side, true, true)
	potentialMoves = append(potentialMoves, potentialUpMoves...)
	potentialDownMoves := getStraightLineMoves(coord, board, side, true, false)
	potentialMoves = append(potentialMoves, potentialDownMoves...)
	potentialRightMoves := getStraightLineMoves(coord, board, side, false, true)
	potentialMoves = append(potentialMoves, potentialRightMoves...)
	potentialLeftMoves := getStraightLineMoves(coord, board, side, false, false)
	potentialMoves = append(potentialMoves, potentialLeftMoves...)
	return potentialMoves
}

// Gets straight line moves in single direction for a given coordinate, board, and side. moveVertical specifies whether piece should
// move vertically or horizontally; increment specifies whether piece should move up or down (if vertical) or left or right (if horzontal)
func getStraightLineMoves(coord Coordinate, board Board, side Side, moveVertical bool, increment bool) []Coordinate {
	var potentialMoves []Coordinate
	var currentChangeVal int
	if increment {
		currentChangeVal = 1
	} else {
		currentChangeVal = -1
	}
	for {
		newCoord := getNextCoordinate(coord, currentChangeVal, moveVertical)
		toAdd, toBreak := canMoveToSquare(newCoord, board, side)
		if toAdd {
			potentialMoves = append(potentialMoves, newCoord)
		}
		if toBreak {
			break
		}
		if increment {
			currentChangeVal++
		} else {
			currentChangeVal--
		}
	}
	return potentialMoves
}

// Returns whether to add coordinate to potential moves list, and whether loop encompassing this method should break (if path stops)
func canMoveToSquare(coord Coordinate, board Board, side Side) (bool, bool) {
	if !coord.isLegal() {
		return false, true
	} else if board.isSpaceOccupied(coord) && board.getPieceSide(coord) == side {
		return false, true
	} else if board.isSpaceOccupied(coord) && board.getPieceSide(coord) == side {
		return true, true
	} else {
		return true, false
	}
}

func getNextCoordinate(coord Coordinate, changeVal int, moveVertical bool) Coordinate {
	if moveVertical {
		newRow := coord.Row + changeVal
		return Coordinate{Row: coord.Row, Column: newRow}
	}
	newCol := coord.Column + changeVal
	return Coordinate{Row: coord.Row, Column: newCol}
}

func (coord Coordinate) isLegal() bool {
	return coord.Row <= 7 && coord.Row >= 0 && coord.Column <= 7 && coord.Column >= 0
}
