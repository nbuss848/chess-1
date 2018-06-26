// moveutil
package chessgame

// Gets all straight line moves, given a piece's coordinates, its side, and a board. Used for rooks and queens
func getAllStraightLineMoves(coord Coordinate, board *ChessBoard, side Side) []Coordinate {
	var allPotentialMoves []Coordinate
	potentialUpMoves := getStraightLineMoves(coord, board, side, true, true)
	allPotentialMoves = append(allPotentialMoves, potentialUpMoves...)

	potentialDownMoves := getStraightLineMoves(coord, board, side, true, false)
	allPotentialMoves = append(allPotentialMoves, potentialDownMoves...)

	potentialRightMoves := getStraightLineMoves(coord, board, side, false, true)
	allPotentialMoves = append(allPotentialMoves, potentialRightMoves...)

	potentialLeftMoves := getStraightLineMoves(coord, board, side, false, false)
	allPotentialMoves = append(allPotentialMoves, potentialLeftMoves...)

	return allPotentialMoves
}

func getAllDiagonalMoves(coord Coordinate, board *ChessBoard, side Side) []Coordinate {
	var allPotentialMoves []Coordinate
	potentialLeftAndUpMoves := getDiagonalMoves(coord, board, side, true, false)
	allPotentialMoves = append(allPotentialMoves, potentialLeftAndUpMoves...)

	potentialRightAndUpMoves := getDiagonalMoves(coord, board, side, true, true)
	allPotentialMoves = append(allPotentialMoves, potentialRightAndUpMoves...)

	potentialLeftAndDownMoves := getDiagonalMoves(coord, board, side, false, false)
	allPotentialMoves = append(allPotentialMoves, potentialLeftAndDownMoves...)

	potentialRightAndDownMoves := getDiagonalMoves(coord, board, side, false, true)
	allPotentialMoves = append(allPotentialMoves, potentialRightAndDownMoves...)

	return allPotentialMoves
}

// Gets straight line moves in single direction for a given coordinate, board, and side. moveVertical specifies whether piece should
// move vertically or horizontally; increment specifies whether piece should move up or down (if vertical) or left or right (if horzontal)
func getStraightLineMoves(coord Coordinate, board *ChessBoard, side Side, moveVertical bool, increment bool) []Coordinate {
	var potentialMoves []Coordinate
	var currentChangeVal int
	if increment {
		currentChangeVal = 1
	} else {
		currentChangeVal = -1
	}
	for {
		newCoord := getNextStraightLineCoordinate(coord, currentChangeVal, moveVertical)
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

func getDiagonalMoves(coord Coordinate, board *ChessBoard, side Side, moveUp bool, moveRight bool) []Coordinate {
	var potentialMoves []Coordinate
	columnChange := 1
	if moveRight {
		columnChange = -1
	}
	rowChange := 1
	if moveUp {
		rowChange = -1
	}
	for {
		newCoord := getNextDiagonalCoordinate(coord, rowChange, columnChange)
		toAdd, toBreak := canMoveToSquare(newCoord, board, side)
		if toAdd {
			potentialMoves = append(potentialMoves, newCoord)
		}
		if toBreak {
			break
		}

		if moveUp {
			rowChange--
		} else {
			rowChange++
		}
		if moveRight {
			columnChange++
		} else {
			columnChange--
		}
	}
	return potentialMoves
}

// Returns whether to add coordinate to potential moves list, and whether loop encompassing this method should break (if path stops)
func canMoveToSquare(coord Coordinate, board *ChessBoard, side Side) (bool, bool) {
	if !coord.isLegal() {
		return false, true
	} else if board.isSpaceOccupied(coord) && board.getPieceSide(coord) == side {
		return false, true
	} else if board.isSpaceOccupied(coord) && board.getPieceSide(coord) != side {
		return true, true
	} else {
		return true, false
	}
}

func getNextStraightLineCoordinate(coord Coordinate, changeVal int, moveVertical bool) Coordinate {
	if moveVertical {
		newRow := coord.Row + changeVal
		return Coordinate{Row: newRow, Column: coord.Column}
	}
	newCol := coord.Column + changeVal
	return Coordinate{Row: coord.Row, Column: newCol}
}

func getNextDiagonalCoordinate(coord Coordinate, verticalChange int, horizontalChange int) Coordinate {
	newRow := coord.Row + verticalChange
	newCol := coord.Column + horizontalChange
	return Coordinate{Row: newRow, Column: newCol}
}

func (coord Coordinate) isLegal() bool {
	return coord.Row <= 7 && coord.Row >= 0 && coord.Column <= 7 && coord.Column >= 0
}

func AbsIntVal(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}
