// chessboard
package chessgame

// Struct representing chess board
type ChessBoard struct {
	BoardPieces [][]ChessPiece
	MoveLog     []Move
	WhiteKing   *King
	BlackKing   *King
}

// Represents a move in a chess game
type Move struct {
	fromCoordinate Coordinate
	toCoordinate   Coordinate
	isCapture      bool
	piece          ChessPiece
}

// Creates new ChessBoard with all pieces in place for beginning of game
func CreateBoard() ChessBoard {
	moveLog := []Move{}
	whiteBaseRow := createBaseRow(0, WHITE)
	whitePawns := createPawnRow(1, WHITE)
	emptyRow1 := createEmptyRow()
	emptyRow2 := createEmptyRow()
	emptyRow3 := createEmptyRow()
	emptyRow4 := createEmptyRow()
	blackPawns := createPawnRow(6, BLACK)
	blackBaseRow := createBaseRow(7, BLACK)
	board := [][]ChessPiece{whiteBaseRow, whitePawns, emptyRow1, emptyRow2, emptyRow3, emptyRow4, blackPawns, blackBaseRow}
	whiteKing := board[0][4].(*King)
	blackKing := board[7][4].(*King)
	return ChessBoard{BoardPieces: board, MoveLog: moveLog, WhiteKing: whiteKing, BlackKing: blackKing}
}

// Creates base row for given side, with rooks, knights, bishops, king and queen
func createBaseRow(row int, pieceSide Side) []ChessPiece {
	leftRook := newRook(pieceSide, Coordinate{Row: row, Column: 0})
	leftKnight := newKnight(pieceSide, Coordinate{Row: row, Column: 1})
	leftbishop := newBishop(pieceSide, Coordinate{Row: row, Column: 2})
	queen := newQueen(pieceSide, Coordinate{Row: row, Column: 3})
	king := newKing(pieceSide, Coordinate{Row: row, Column: 4})
	rightBishop := newBishop(pieceSide, Coordinate{Row: row, Column: 5})
	rightKnight := newKnight(pieceSide, Coordinate{Row: row, Column: 6})
	rightRook := newRook(pieceSide, Coordinate{Row: row, Column: 7})
	return []ChessPiece{&leftRook, &leftKnight, &leftbishop, &queen, &king, &rightBishop, &rightKnight, &rightRook}
}

// Creates row of pawns for ChessBoard, all of given side
func createPawnRow(row int, pieceSide Side) []ChessPiece {
	var pawns []ChessPiece
	for i := 0; i < 8; i++ {
		pawn := newPawn(pieceSide, Coordinate{Row: row, Column: i})
		pawns = append(pawns, &pawn)
	}
	return pawns
}

// Creates empty row for ChessBoard
func createEmptyRow() []ChessPiece {
	return []ChessPiece{nil, nil, nil, nil, nil, nil, nil, nil}
}

// Creates empty ChessBoard - for testing purposes
func createEmptyBoard() ChessBoard {
	emptyBoard := [][]ChessPiece{}
	for i := 0; i < 8; i++ {
		emptyBoard = append(emptyBoard, createEmptyRow())
	}
	moveLog := []Move{}
	// Assign kings fake coordinates for testing purposes
	whiteKing := newKing(WHITE, Coordinate{-1, -1})
	blackKing := newKing(BLACK, Coordinate{8, 8})
	return ChessBoard{BoardPieces: emptyBoard, MoveLog: moveLog, WhiteKing: &whiteKing, BlackKing: &blackKing}
}

// Checks whether given space on board is occupied
func (board *ChessBoard) isSpaceOccupied(coord Coordinate) bool {
	if board.BoardPieces[coord.Row][coord.Column] != nil {
		return true
	}
	return false
}

// Gets type of piece occupying given coordinate
func (board *ChessBoard) getPieceType(coord Coordinate) PieceType {
	return board.BoardPieces[coord.Row][coord.Column].getPieceType()
}

// Gets side of piece occupying given coordinate
func (board *ChessBoard) getPieceSide(coord Coordinate) Side {
	return board.BoardPieces[coord.Row][coord.Column].getPieceSide()
}

// Gets previous move on board
func (board *ChessBoard) getPreviousMove() (Move, bool) {
	if len(board.MoveLog) == 0 {
		return Move{}, false
	}
	return board.MoveLog[len(board.MoveLog)-1], true
}

// Checks whether piece occupying given coordinate has moved
func (board *ChessBoard) hasPieceMoved(coord Coordinate) bool {
	return board.BoardPieces[coord.Row][coord.Column].hasPieceMoved()
}

// Returns Coordinate of king for given side
func (board *ChessBoard) getKingCoordinate(pieceSide Side) Coordinate {
	if pieceSide == WHITE {
		return board.WhiteKing.currentCoordinate
	}
	return board.BlackKing.currentCoordinate
}

// Returns all valid moves for given side, as a dictionary with coordinates as keys and maps of coordinates and booleans as the values
func (board *ChessBoard) getAllValidMovesForSide(pieceSide Side) map[Coordinate]map[Coordinate]bool {
	moves := make(map[Coordinate]map[Coordinate]bool)
	for row := 0; row < len(board.BoardPieces); row++ {
		for col := 0; col < len(board.BoardPieces[row]); col++ {
			currentCoord := Coordinate{row, col}
			if board.isSpaceOccupied(currentCoord) && board.getPieceSide(currentCoord) == pieceSide {
				moves[currentCoord] = board.BoardPieces[row][col].validMoves(board)
			}
		}
	}
	return moves
}
