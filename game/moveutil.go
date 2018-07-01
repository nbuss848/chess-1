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
// move vertically or horizontally; increment specifies whether piece should move up or down (if vertical) or left or right (if horizontal)
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
	columnChange := -1
	if moveRight {
		columnChange = 1
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

func getAllMovesForPiece(board *ChessBoard, piece ChessPiece, allMoves func(*ChessBoard, ChessPiece) map[Coordinate]bool) map[Coordinate]bool {
	kingCoord := board.getKingCoordinate(piece.getPieceSide())
	if willMoveExposeKing(kingCoord, piece.getCurrentCoordinates(), piece.getPieceSide(), board) {
		return nil
	}
	validMoves := allMoves(board, piece)
	king := board.WhiteKing
	if piece.getPieceSide() == BLACK {
		king = board.BlackKing
	}
	if !king.inCheck {
		return validMoves
	}
	if len(king.threateningPieces) > 1 {
		return nil
	}
	validCheckMoves := make(map[Coordinate]bool)
	_, ok := validMoves[king.threateningPieces[0].getCurrentCoordinates()]
	if ok {
		validCheckMoves[king.threateningPieces[0].getCurrentCoordinates()] = true
	}
	checkPieceType := king.threateningPieces[0].getPieceType()
	checkPieceCoord := king.threateningPieces[0].getCurrentCoordinates()
	blockingCoordinates := getCheckBlockingCoords(kingCoord, checkPieceCoord, validMoves, checkPieceType)
	for i := 0; i < len(blockingCoordinates); i++ {
		validCheckMoves[blockingCoordinates[i]] = true
	}
	if len(validCheckMoves) == 0 {
		return nil
	}
	return validCheckMoves
}

// scans all coordinates between king and piece that is putting king in check, and returns
// scanned coordinates that are in possibleMoves
func getCheckBlockingCoords(kingCoord Coordinate, checkPieceCoord Coordinate, possibleMoves map[Coordinate]bool, checkPieceType PieceType) []Coordinate {
	var blockingCoordinates []Coordinate
	if checkPieceType == PAWN || checkPieceType == KNIGHT {
		return blockingCoordinates
	}
	rowMove := 0
	if kingCoord.Row > checkPieceCoord.Row {
		rowMove = -1
	} else if kingCoord.Row < checkPieceCoord.Row {
		rowMove = 1
	}
	colMove := 0
	if kingCoord.Column > checkPieceCoord.Column {
		colMove = -1
	} else if kingCoord.Column < checkPieceCoord.Column {
		colMove = 1
	}
	currentCoord := Coordinate{kingCoord.Row + rowMove, kingCoord.Column + colMove}
	for currentCoord != checkPieceCoord {
		_, ok := possibleMoves[currentCoord]
		if ok {
			blockingCoordinates = append(blockingCoordinates, currentCoord)
		}
		currentCoord.Row += rowMove
		currentCoord.Column += colMove
	}
	return blockingCoordinates
}

// Checks whether two coordinates are aligned across straight lines and diagonals
func areCoordinatesAligned(coord1 Coordinate, coord2 Coordinate) bool {
	if coord1.Row == coord2.Row {
		return true
	}
	if coord1.Column == coord2.Column {
		return true
	}
	rowAbsDif := AbsIntVal(coord1.Row - coord2.Row)
	colAbsDif := AbsIntVal(coord1.Column - coord2.Column)
	if rowAbsDif == colAbsDif {
		return true
	}
	return false
}
