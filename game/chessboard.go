// chessboard
package chessgame

type ChessBoard struct {
	BoardPieces [][]ChessPiece
	MoveLog     []Move
}

type Move struct {
	fromCoordinate Coordinate
	toCoordinate   Coordinate
	isCapture      bool
	piece          ChessPiece
}

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
	return ChessBoard{BoardPieces: board, MoveLog: moveLog}
}

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

func createPawnRow(row int, pieceSide Side) []ChessPiece {
	var pawns []ChessPiece
	for i := 0; i < 8; i++ {
		pawn := newPawn(pieceSide, Coordinate{Row: row, Column: i})
		pawns = append(pawns, &pawn)
	}
	return pawns
}

func createEmptyRow() []ChessPiece {
	return []ChessPiece{nil, nil, nil, nil, nil, nil, nil, nil}
}

func createEmptyBoard() ChessBoard {
	emptyBoard := [][]ChessPiece{}
	for i := 0; i < 8; i++ {
		emptyBoard = append(emptyBoard, createEmptyRow())
	}
	moveLog := []Move{}
	return ChessBoard{BoardPieces: emptyBoard, MoveLog: moveLog}
}

// TODO: combine isSpaceOccupied with getPieceSide
func (board *ChessBoard) isSpaceOccupied(coord Coordinate) bool {
	if board.BoardPieces[coord.Row][coord.Column] != nil {
		return true
	}
	return false
}

func (board *ChessBoard) getPieceType(coord Coordinate) PieceType {
	return board.BoardPieces[coord.Row][coord.Column].getPieceType()
}

func (board *ChessBoard) getPieceSide(coord Coordinate) Side {
	return board.BoardPieces[coord.Row][coord.Column].getPieceSide()
}

func (board *ChessBoard) getPreviousMove() (Move, bool) {
	if len(board.MoveLog) == 0 {
		return Move{}, false
	}
	return board.MoveLog[len(board.MoveLog)-1], true
}

func (board *ChessBoard) hasPieceMoved(coord Coordinate) bool {
	return board.BoardPieces[coord.Row][coord.Column].hasPieceMoved()
}
