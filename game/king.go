// king
package chessgame

type King struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	potentialMoves    map[Coordinate]bool
	inCheck           bool
	canCastle         bool
}

func (king *King) updatePosition(newCoord Coordinate) {
	king.currentCoordinate = newCoord
	king.canCastle = false
}

func (king *King) updateValidMoves(board Board) {
	king.potentialMoves = make(map[Coordinate]bool)
	potentialCoordinates := getSurroundingCoordinates(king.currentCoordinate)
	for i := 0; i < len(potentialCoordinates); i++ {
		if !willKingMoveLeadToCheck(potentialCoordinates[i], board, king.pieceSide) {
			king.potentialMoves[potentialCoordinates[i]] = true
		}
	}
	//TODO add support for castling
}

func willKingMoveLeadToCheck(coord Coordinate, board Board, pieceSide Side) bool {
	if isSpaceThreatenedByPawn(coord, board, pieceSide) {
		return true
	}
	if isSpaceThreatenedByKnight(coord, board, pieceSide) {
		return true
	}
	if isSpaceThreatenedByAnyDiagonals(coord, board, pieceSide) {
		return true
	}
	if isSpaceThreatenedByAnyStraightLines(coord, board, pieceSide) {
		return true
	}
	if isSpaceThreatenedByKing(coord, board, pieceSide) {
		return true
	}
	return false
}

func isSpaceThreatenedByPawn(coord Coordinate, board Board, pieceSide Side) bool {
	firstPawnCol := coord.Column + 1
	secondPawnCol := coord.Column - 1
	threateningRow := coord.Row - 1
	if pieceSide == BLACK {
		threateningRow = coord.Row - 2
	}
	firstPawnCoord := Coordinate{Row: threateningRow, Column: firstPawnCol}
	secondPawnCoord := Coordinate{Row: threateningRow, Column: secondPawnCol}
	if canCoordinateThreaten(board, firstPawnCoord, pieceSide, PAWN) {
		return true
	}
	if canCoordinateThreaten(board, secondPawnCoord, pieceSide, PAWN) {
		return true
	}
	return false
}

func isSpaceThreatenedByKnight(coord Coordinate, board Board, pieceSide Side) bool {
	possibleKnightPositions := getAllPossibleKnightMoves(coord)
	for i := 0; i < len(possibleKnightPositions); i++ {
		if canCoordinateThreaten(board, possibleKnightPositions[i], pieceSide, KNIGHT) {
			return true
		}
	}
	return false
}

func isSpaceThreatenedByAnyDiagonals(coord Coordinate, board Board, pieceSide Side) bool {
	if isSpaceThreatenedByDiagonal(coord, board, pieceSide, -1, 1) {
		return true
	}
	if isSpaceThreatenedByDiagonal(coord, board, pieceSide, -1, -1) {
		return true
	}
	if isSpaceThreatenedByDiagonal(coord, board, pieceSide, 1, -1) {
		return true
	}
	if isSpaceThreatenedByDiagonal(coord, board, pieceSide, 1, 1) {
		return true
	}
	return false
}

func isSpaceThreatenedByAnyStraightLines(coord Coordinate, board Board, pieceSide Side) bool {
	if isSpaceThreatenedByStraightLine(board, coord, pieceSide, false, -1) {
		return true
	}
	if isSpaceThreatenedByStraightLine(board, coord, pieceSide, false, 1) {
		return true
	}
	if isSpaceThreatenedByStraightLine(board, coord, pieceSide, true, -1) {
		return true
	}
	if isSpaceThreatenedByStraightLine(board, coord, pieceSide, true, 1) {
		return true
	}
	return false
}

func isSpaceThreatenedByKing(coord Coordinate, board Board, pieceSide Side) bool {
	potentialCoords := getSurroundingCoordinates(coord)
	for i := 0; i < len(potentialCoords); i++ {
		if canCoordinateThreaten(board, coord, pieceSide, KING) {
			return true
		}
	}
	return false
}

func isSpaceThreatenedByStraightLine(board Board, coord Coordinate, pieceSide Side, vertical bool, changeVal int) bool {
	newRow := coord.Row
	newColumn := coord.Column
	if vertical {
		newRow += changeVal
	} else {
		newColumn += changeVal
	}
	currentCoord := Coordinate{Row: newRow, Column: newColumn}
	for {
		if !currentCoord.isLegal() {
			return false
		}
		if board.isSpaceOccupied(coord) {
			canThreatenQueen := canCoordinateThreaten(board, currentCoord, pieceSide, QUEEN)
			canThreatenRook := canCoordinateThreaten(board, currentCoord, pieceSide, ROOK)
			if canThreatenQueen || canThreatenRook {
				return true
			}
			return false
		}
		if vertical {
			currentCoord.Row += changeVal
		} else {
			currentCoord.Column += changeVal
		}
	}
	return false
}

func isSpaceThreatenedByDiagonal(coord Coordinate, board Board, pieceSide Side, rowChange int, colChange int) bool {
	currentCoord := Coordinate{Row: coord.Row + rowChange, Column: coord.Column + colChange}
	for {
		if !currentCoord.isLegal() {
			return false
		}
		if board.isSpaceOccupied(currentCoord) {
			canThreatenQueen := canCoordinateThreaten(board, currentCoord, pieceSide, QUEEN)
			canThreatenBishop := canCoordinateThreaten(board, currentCoord, pieceSide, BISHOP)
			if canThreatenBishop || canThreatenQueen {
				return true
			}
			return false
		}
		currentCoord.Column += colChange
		currentCoord.Row += rowChange
	}
	return false
}

func canCoordinateThreaten(board Board, coord Coordinate, pieceSide Side, pieceType PieceType) bool {
	if !coord.isLegal() {
		return false
	}
	if !board.isSpaceOccupied(coord) {
		return false
	}
	if board.getPieceSide(coord) == pieceSide {
		return false
	}
	if board.getPieceType(coord) != pieceType {
		return false
	}
	return true
}

func getSurroundingCoordinates(coord Coordinate) []Coordinate {
	upperCol := coord.Column + 1
	lowerCol := coord.Column - 1
	upperRow := coord.Row + 1
	lowerRow := coord.Row - 1

	var coords []Coordinate
	for currentCol := lowerCol; currentCol <= upperCol; currentCol++ {
		coords = append(coords, Coordinate{Row: upperRow, Column: currentCol})
		coords = append(coords, Coordinate{Row: lowerRow, Column: currentCol})
	}
	coords = append(coords, Coordinate{Row: coord.Row, Column: lowerCol})
	coords = append(coords, Coordinate{Row: coord.Row, Column: upperCol})
	return coords
}
