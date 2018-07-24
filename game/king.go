// king
package chessgame

// Struct representation of King piece
type King struct {
	id                int
	currentCoordinate Coordinate
	pieceSide         Side
	inCheck           bool
	hasMoved          bool
	threateningPieces []ThreateningPiece
}

// Sets hasMoved property in king struct. Used for cloning
func (king *King) setHasMoved(hasMoved bool) {
	king.hasMoved = hasMoved
}

// Returns King's current coordinate
func (king *King) GetCurrentCoordinates() Coordinate {
	return king.currentCoordinate
}

// Creates new king
func newKing(kingSide Side, coord Coordinate) King {
	return King{currentCoordinate: coord, pieceSide: kingSide, hasMoved: false, inCheck: false}
}

// Update's king's position
func (king *King) updatePosition(newCoord Coordinate) {
	king.currentCoordinate = newCoord
	king.hasMoved = true
}

// Checks whether King has moved
func (king *King) hasPieceMoved() bool {
	return king.hasMoved
}

// Gets all valid moves for King
func (king *King) ValidMoves(board *ChessBoard) map[Coordinate]bool {
	validMoves := make(map[Coordinate]bool)
	potentialCoordinates := getSurroundingCoordinates(king.currentCoordinate)
	// temporarily set current king space to nil, to simulate board if king were to move
	board.BoardPieces[king.currentCoordinate.Row][king.currentCoordinate.Column] = nil
	for i := 0; i < len(potentialCoordinates); i++ {
		if potentialCoordinates[i].isLegal() && !isSpaceOccupiedBySameSidePiece(potentialCoordinates[i], board, king.pieceSide) && !willKingMoveLeadToCheck(potentialCoordinates[i], board, king.pieceSide) {
			validMoves[potentialCoordinates[i]] = true
		}
	}
	board.BoardPieces[king.currentCoordinate.Row][king.currentCoordinate.Column] = king
	if king.canCastle(board, true) {
		validMoves[getCastleCoordinate(king.currentCoordinate, true)] = true
	}
	if king.canCastle(board, false) {
		validMoves[getCastleCoordinate(king.currentCoordinate, false)] = true
	}
	return validMoves
}

// Checks if coordinate is occupied by piece of given side
func isSpaceOccupiedBySameSidePiece(coord Coordinate, board *ChessBoard, pieceSide Side) bool {
	if !board.isSpaceOccupied(coord) {
		return false
	}
	if board.GetPieceSide(coord) != pieceSide {
		return false
	}
	return true
}

// Return's king's piece type
func (king *King) GetPieceType() PieceType {
	return KING
}

// Returns king's piece side
func (king *King) GetPieceSide() Side {
	return king.pieceSide
}

// Returns boolean for whether king can castle
func (king *King) canCastle(board *ChessBoard, castleLeft bool) bool {
	if king.hasMoved {
		return false
	}
	if king.inCheck {
		return false
	}
	changeVal := 1
	if castleLeft {
		changeVal = -1
	}
	boardEdge := 7
	if castleLeft {
		boardEdge = 0
	}
	newCoord := Coordinate{Row: king.currentCoordinate.Row, Column: king.currentCoordinate.Column + changeVal}
	for newCoord.isLegal() {
		if board.isSpaceOccupied(newCoord) && newCoord.Column != boardEdge {
			return false
		}
		if newCoord.Column == boardEdge && !board.isSpaceOccupied(newCoord) {
			return false
		}
		if newCoord.Column == boardEdge && (board.GetPieceType(newCoord) == ROOK && board.GetPieceSide(newCoord) == king.pieceSide && !board.BoardPieces[newCoord.Row][newCoord.Column].hasPieceMoved()) {
			return true
		}
		newCoord.Column += changeVal
	}
	return false
}

// Gets coordinate King would move to if a castle took place
func getCastleCoordinate(coord Coordinate, castleLeft bool) Coordinate {
	if castleLeft {
		return Coordinate{Row: coord.Row, Column: coord.Column - 2}
	} else {
		return Coordinate{Row: coord.Row, Column: coord.Column + 2}
	}
}

// Checks if potential move by king will lead to check
func willKingMoveLeadToCheck(coord Coordinate, board *ChessBoard, pieceSide Side) bool {
	if len(threateningPawnCoordinates(coord, board, pieceSide)) > 0 {
		return true
	}
	if len(threateningKnightCoordinates(coord, board, pieceSide)) > 0 {
		return true
	}
	if len(threateningDiagonalCoords(coord, board, pieceSide)) > 0 {
		return true
	}
	if len(threateningStraightLineCoords(coord, board, pieceSide)) > 0 {
		return true
	}
	if isSpaceThreatenedByKing(coord, board, pieceSide) {
		return true
	}
	return false
}

// Checks if space is threatened by pawn
func threateningPawnCoordinates(coord Coordinate, board *ChessBoard, pieceSide Side) []Coordinate {
	var threateningCoordinates []Coordinate
	firstPawnCol := coord.Column + 1
	secondPawnCol := coord.Column - 1
	threateningRow := coord.Row + 1
	if pieceSide == BLACK {
		threateningRow = coord.Row - 1
	}
	firstPawnCoord := Coordinate{Row: threateningRow, Column: firstPawnCol}
	secondPawnCoord := Coordinate{Row: threateningRow, Column: secondPawnCol}
	if canCoordinateThreaten(board, firstPawnCoord, pieceSide, PAWN) {
		threateningCoordinates = append(threateningCoordinates, firstPawnCoord)
	}
	if canCoordinateThreaten(board, secondPawnCoord, pieceSide, PAWN) {
		threateningCoordinates = append(threateningCoordinates, secondPawnCoord)
	}
	return threateningCoordinates
}

// Checks if space is threatened by a knight
func threateningKnightCoordinates(coord Coordinate, board *ChessBoard, pieceSide Side) []Coordinate {
	var threateningCoordinates []Coordinate
	possibleKnightPositions := getAllPossibleKnightMoves(coord)
	for i := 0; i < len(possibleKnightPositions); i++ {
		if canCoordinateThreaten(board, possibleKnightPositions[i], pieceSide, KNIGHT) {
			threateningCoordinates = append(threateningCoordinates, possibleKnightPositions[i])
		}
	}
	return threateningCoordinates
}

// Checks whether space is threatened by opposing King
func isSpaceThreatenedByKing(coord Coordinate, board *ChessBoard, pieceSide Side) bool {
	potentialCoords := getSurroundingCoordinates(coord)
	for i := 0; i < len(potentialCoords); i++ {
		if canCoordinateThreaten(board, coord, pieceSide, KING) {
			return true
		}
	}
	return false
}

// Checks if space is threatened by any straight lines
func threateningStraightLineCoords(coord Coordinate, board *ChessBoard, pieceSide Side) []Coordinate {
	var threateningCoordinates []Coordinate
	threat := isSpaceThreatenedAcrossLine(coord, board, pieceSide, ROOK)

	left := iterateCoordinates(coord, board, pieceSide, 0, -1, threat)
	threateningCoordinates = append(threateningCoordinates, left...)

	right := iterateCoordinates(coord, board, pieceSide, 0, 1, threat)
	threateningCoordinates = append(threateningCoordinates, right...)

	up := iterateCoordinates(coord, board, pieceSide, -1, 0, threat)
	threateningCoordinates = append(threateningCoordinates, up...)

	down := iterateCoordinates(coord, board, pieceSide, 1, 0, threat)
	threateningCoordinates = append(threateningCoordinates, down...)
	return threateningCoordinates
}

// Checks if space is threatened by any diagonals
func threateningDiagonalCoords(coord Coordinate, board *ChessBoard, pieceSide Side) []Coordinate {
	var threateningCoordinates []Coordinate
	threat := isSpaceThreatenedAcrossLine(coord, board, pieceSide, BISHOP)

	leftUp := iterateCoordinates(coord, board, pieceSide, -1, -1, threat)
	threateningCoordinates = append(threateningCoordinates, leftUp...)

	leftDown := iterateCoordinates(coord, board, pieceSide, 1, -1, threat)
	threateningCoordinates = append(threateningCoordinates, leftDown...)

	rightUp := iterateCoordinates(coord, board, pieceSide, -1, 1, threat)
	threateningCoordinates = append(threateningCoordinates, rightUp...)

	rightDown := iterateCoordinates(coord, board, pieceSide, 1, 1, threat)
	threateningCoordinates = append(threateningCoordinates, rightDown...)
	return threateningCoordinates
}

// Returns a function that uses piece type and can be passed into iterateCoordinates function
func isSpaceThreatenedAcrossLine(coord Coordinate, board *ChessBoard, pieceSide Side, pieceType PieceType) func(Coordinate, *ChessBoard, Side) (bool, bool) {
	return func(coord Coordinate, board *ChessBoard, pieceSide Side) (bool, bool) {
		if !board.isSpaceOccupied(coord) {
			return false, false
		}
		if board.GetPieceSide(coord) == pieceSide {
			return false, true
		}
		coordPieceType := board.GetPieceType(coord)
		if coordPieceType != pieceType && coordPieceType != QUEEN {
			return false, true
		}
		return true, true
	}
}

// Given a board, a coordinate, a side, and a piece type, returns whether coordinate can threaten king
func canCoordinateThreaten(board *ChessBoard, coord Coordinate, pieceSide Side, pieceType PieceType) bool {
	if !coord.isLegal() {
		return false
	}
	if !board.isSpaceOccupied(coord) {
		return false
	}
	if board.GetPieceSide(coord) == pieceSide {
		return false
	}
	if board.GetPieceType(coord) != pieceType {
		return false
	}
	return true
}

// Gets all coordinates surrounding given coordinate (regardless of whether they are valid)
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

// Checks whether moving given piece (specified by pieceToMoveCoord) will expose the king
func willMoveExposeKing(kingCoord Coordinate, pieceToMoveCoord Coordinate, pieceSide Side, board *ChessBoard) (Coordinate, bool) {
	if !areCoordinatesAligned(kingCoord, pieceToMoveCoord) {
		return Coordinate{}, false
	}
	return doesPathContainThreat(board, kingCoord, pieceToMoveCoord, pieceSide)
}

// Checks whether moving the piece at pieceToMoveCoord will create a threat to the king
func doesPathContainThreat(board *ChessBoard, kingCoord Coordinate, pieceToMoveCoord Coordinate, pieceSide Side) (Coordinate, bool) {
	rowChange := 0
	if pieceToMoveCoord.Row > kingCoord.Row {
		rowChange = 1
	} else if pieceToMoveCoord.Row < kingCoord.Row {
		rowChange = -1
	}
	colChange := 0
	if pieceToMoveCoord.Column > kingCoord.Column {
		colChange = 1
	} else if pieceToMoveCoord.Column < kingCoord.Column {
		colChange = -1
	}
	rookOrBishopThreat := ROOK
	if kingCoord.Row != pieceToMoveCoord.Row && kingCoord.Column != pieceToMoveCoord.Column {
		rookOrBishopThreat = BISHOP
	}
	currentCoord := Coordinate{kingCoord.Row + rowChange, kingCoord.Column + colChange}
	pastPieceToMove := false
	for currentCoord.isLegal() {
		if currentCoord == pieceToMoveCoord {
			pastPieceToMove = true
			currentCoord.Column += colChange
			currentCoord.Row += rowChange
			continue
		}
		if !board.isSpaceOccupied(currentCoord) {
			currentCoord.Column += colChange
			currentCoord.Row += rowChange
			continue
		}
		if !pastPieceToMove {
			return Coordinate{}, false
		}
		if pastPieceToMove && board.GetPieceSide(currentCoord) == pieceSide {
			return Coordinate{}, false
		}
		currentCoordPieceType := board.GetPieceType(currentCoord)
		if currentCoordPieceType == QUEEN || currentCoordPieceType == rookOrBishopThreat {
			return currentCoord, true
		}
		return Coordinate{}, false
	}
	return Coordinate{}, false
}

func GetThreateningCoordinates(board *ChessBoard, coords Coordinate, side Side) []Coordinate {
	var threateningCoords []Coordinate
	threateningCoords = append(threateningCoords, threateningStraightLineCoords(coords, board, side)...)
	threateningCoords = append(threateningCoords, threateningDiagonalCoords(coords, board, side)...)
	threateningCoords = append(threateningCoords, threateningKnightCoordinates(coords, board, side)...)
	threateningCoords = append(threateningCoords, threateningPawnCoordinates(coords, board, side)...)
	return threateningCoords
}

// Updates king's status - whether king is in check and what pieces are threatening king if it is in check
func (king *King) updateKingStatus(board *ChessBoard) {
	threateningCoords := GetThreateningCoordinates(board, king.GetCurrentCoordinates(), king.GetPieceSide())
	if len(threateningCoords) == 0 {
		king.inCheck = false
		return
	}
	king.inCheck = true
	var threateningPieces []ThreateningPiece
	for i := 0; i < len(threateningCoords); i++ {
		coord := threateningCoords[i]
		threateningPieces = append(threateningPieces, ThreateningPiece{coord, board.BoardPieces[coord.Row][coord.Column].GetPieceType()})
	}
	king.threateningPieces = threateningPieces
}
