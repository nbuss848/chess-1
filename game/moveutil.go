// moveutil
package chessgame

// Gets all straight line moves, given a piece's coordinates, its side, and a board. Used for rooks and queens
func getAllStraightLineMoves(coord Coordinate, board *ChessBoard, side Side) []Coordinate {
	var allPotentialMoves []Coordinate
	potentialUpMoves := iterateCoordinates(coord, board, side, -1, 0, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialUpMoves...)

	potentialDownMoves := iterateCoordinates(coord, board, side, 1, 0, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialDownMoves...)

	potentialRightMoves := iterateCoordinates(coord, board, side, 0, 1, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialRightMoves...)

	potentialLeftMoves := iterateCoordinates(coord, board, side, 0, -1, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialLeftMoves...)

	return allPotentialMoves
}

func getAllDiagonalMoves(coord Coordinate, board *ChessBoard, side Side) []Coordinate {
	var allPotentialMoves []Coordinate
	potentialLeftAndUpMoves := iterateCoordinates(coord, board, side, -1, -1, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialLeftAndUpMoves...)

	potentialRightAndUpMoves := iterateCoordinates(coord, board, side, -1, 1, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialRightAndUpMoves...)

	potentialLeftAndDownMoves := iterateCoordinates(coord, board, side, 1, -1, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialLeftAndDownMoves...)

	potentialRightAndDownMoves := iterateCoordinates(coord, board, side, 1, 1, canMoveToSquare)
	allPotentialMoves = append(allPotentialMoves, potentialRightAndDownMoves...)

	return allPotentialMoves
}

// Iterates through coordinates based on row change and column change, adding coordinates using logic from appendLogic func
func iterateCoordinates(coord Coordinate, board *ChessBoard, side Side, rowChange int, colChange int, appendLogic func(Coordinate, *ChessBoard, Side) (bool, bool)) []Coordinate {
	var coordinates []Coordinate
	currentCoord := getNextCoordinate(coord, rowChange, colChange)
	for currentCoord.isLegal() {
		appendCoord, toBreak := appendLogic(currentCoord, board, side)
		if appendCoord {
			coordinates = append(coordinates, currentCoord)
		}
		if toBreak {
			break
		}
		currentCoord = getNextCoordinate(currentCoord, rowChange, colChange)
	}
	return coordinates
}

// Given an amount to change the row and column by, gets the next coordinate
func getNextCoordinate(coord Coordinate, rowChange int, colChange int) Coordinate {
	return Coordinate{coord.Row + rowChange, coord.Column + colChange}
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

// Checks whether coordinate is within legal bounds of chess board
func (coord Coordinate) isLegal() bool {
	return coord.Row <= 7 && coord.Row >= 0 && coord.Column <= 7 && coord.Column >= 0
}

// Returns absolute value of integer - Go Math library only provides this function for floats
func AbsIntVal(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

// Gets all possible moves for piece, taking king's check status and possibility of exposing king into account
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
