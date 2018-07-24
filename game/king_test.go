// chessboard_test
package chessgame

import (
	"testing"
)

func TestKingMovesMiddleOfBoard(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{4, 4})
	board.BoardPieces[4][4] = &king
	moves := king.ValidMoves(&board)
	if len(moves) != 8 {
		t.Fatalf("Expected 8 moves, got %d", len(moves))
	}
}

func TestKingMovesNoCheck(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	moves := king.ValidMoves(&board)
	if len(moves) != 5 {
		t.Fatalf("Expected 5 moves, got %d instead", len(moves))
	}
}

func TestKingMovesShouldBeEmpty(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	rook1 := newRook(BLACK, Coordinate{1, 5})
	rook2 := newRook(BLACK, Coordinate{1, 3})
	board.BoardPieces[1][5] = &rook1
	board.BoardPieces[1][3] = &rook2
	moves := king.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected 0 moves, got %d instead", len(moves))
	}
}

func TestKingMovesShouldCaptureRook(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	rook1 := newRook(BLACK, Coordinate{1, 5})
	board.BoardPieces[1][5] = &rook1
	moves := king.ValidMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 5 moves, got %d instead", len(moves))
	}
	captureCoord := Coordinate{Row: 1, Column: 5}
	_, ok := moves[captureCoord]
	if !ok {
		t.Fatalf("Row 1, column 5 should be a valid move")
	}
	otherCoord := Coordinate{Row: 0, Column: 3}
	_, ok = moves[otherCoord]
	if !ok {
		t.Fatalf("Row 0, column 3 should be the other valid move")
	}
}

func TestKingMovesWithRookAndQueen(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	rook := newRook(BLACK, Coordinate{1, 5})
	board.BoardPieces[1][5] = &rook
	queen := newQueen(BLACK, Coordinate{2, 5})
	board.BoardPieces[2][5] = &queen
	moves := king.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected no moves, got %d instead", len(moves))
	}
	queen.pieceSide = WHITE
	moves = king.ValidMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 2 moves, got %d instead", len(moves))
	}
}

func TestKingMovesTrappedByOwnPieces(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	rook1 := newRook(WHITE, Coordinate{0, 5})
	rook2 := newRook(WHITE, Coordinate{0, 3})
	rook3 := newRook(WHITE, Coordinate{1, 5})
	rook4 := newRook(WHITE, Coordinate{1, 3})
	rook5 := newRook(WHITE, Coordinate{1, 4})
	board.BoardPieces[0][5] = &rook1
	board.BoardPieces[0][3] = &rook2
	board.BoardPieces[1][5] = &rook3
	board.BoardPieces[1][3] = &rook4
	board.BoardPieces[1][4] = &rook5
	moves := king.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected 0 moves, got %d instead", len(moves))
	}
}

func TestKingMovesCanCastle(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	rook1 := newRook(WHITE, Coordinate{0, 0})
	rook2 := newRook(WHITE, Coordinate{0, 7})
	board.BoardPieces[0][0] = &rook1
	board.BoardPieces[0][7] = &rook2
	moves := king.ValidMoves(&board)
	firstCoord := Coordinate{0, 2}
	_, ok := moves[firstCoord]
	if !ok {
		t.Fatalf("Row 0 column 2 should be a valid move")
	}
	if len(moves) != 7 {
		t.Fatalf("Expected 7 moves, got %d instead", len(moves))
	}
}

func TestMoveExposesKing(t *testing.T) {
	board := createEmptyBoard()
	board.WhiteKing.currentCoordinate = Coordinate{0, 0}
	rook := newRook(WHITE, Coordinate{1, 1})
	board.BoardPieces[1][1] = &rook
	enemyBishop := newBishop(BLACK, Coordinate{2, 2})
	board.BoardPieces[2][2] = &enemyBishop
	moves := rook.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected 0 moves, got %d instead", len(moves))
	}
}

func TestMoveCapturesThreateningPiece(t *testing.T) {
	board := createEmptyBoard()
	board.BlackKing.currentCoordinate = Coordinate{7, 4}
	enemyQueen := newQueen(WHITE, Coordinate{6, 4})
	board.BoardPieces[6][4] = &enemyQueen
	board.BlackKing.inCheck = true
	board.BlackKing.threateningPieces = []ThreateningPiece{ThreateningPiece{enemyQueen.currentCoordinate, enemyQueen.GetPieceType()}}
	knight := newKnight(BLACK, Coordinate{7, 2})
	board.BoardPieces[4][3] = &knight
	moves := knight.ValidMoves(&board)
	if len(moves) != 1 {
		t.Fatalf("Expected 1 move, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{6, 4}]
	if !ok {
		t.Fatalf("Only move should be to row 6 column 4")
	}
}

func TestMoveBlocksThreateningPiece(t *testing.T) {
	board := createEmptyBoard()
	board.BlackKing.currentCoordinate = Coordinate{7, 4}
	enemyQueen := newQueen(WHITE, Coordinate{5, 4})
	board.BoardPieces[6][4] = &enemyQueen
	board.BlackKing.inCheck = true
	board.BlackKing.threateningPieces = []ThreateningPiece{ThreateningPiece{enemyQueen.currentCoordinate, enemyQueen.GetPieceType()}}
	rook := newRook(BLACK, Coordinate{6, 7})
	board.BoardPieces[6][7] = &rook
	moves := rook.ValidMoves(&board)
	if len(moves) != 1 {
		t.Fatalf("Expected 1 move, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{6, 4}]
	if !ok {
		t.Fatalf("Only move should be to row 6 column 4")
	}
	pawn := newPawn(WHITE, Coordinate{6, 6})
	board.BoardPieces[6][6] = &pawn
	moves = rook.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected no moves, due to blocking pawn. Got %d instead", len(moves))
	}
}

func TestMovesKingThreatenedByMultiplePieces(t *testing.T) {
	board := createEmptyBoard()
	board.BlackKing.currentCoordinate = Coordinate{7, 4}
	enemyQueen := newQueen(WHITE, Coordinate{5, 4})
	board.BoardPieces[5][4] = &enemyQueen
	board.BoardPieces[7][4] = board.BlackKing
	enemyKnight := newKnight(WHITE, Coordinate{5, 5})
	board.BlackKing.inCheck = true
	board.BlackKing.threateningPieces = []ThreateningPiece{ThreateningPiece{enemyQueen.currentCoordinate, enemyQueen.GetPieceType()}, ThreateningPiece{enemyKnight.currentCoordinate, enemyKnight.GetPieceType()}}
	queen := newQueen(BLACK, Coordinate{6, 5})
	board.BoardPieces[6][5] = &queen
	moves := queen.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected no moves, got %d instead", len(moves))
	}
	moves = board.BlackKing.ValidMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 2 valid moves for king, got %d instead", len(moves))
	}
}

func TestMovesKingThreatenedByPawn(t *testing.T) {
	board := CreateBoard()
	board.BoardPieces[1][3] = nil
	blackPawn := newPawn(BLACK, Coordinate{1, 4})
	board.BoardPieces[1][4] = &blackPawn
	board.WhiteKing.updateKingStatus(&board)
	if !board.WhiteKing.inCheck {
		t.Fatalf("King should be in check")
	}
}

func TestMovesKingLetsBlockingPieceCapture(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 0})
	board.WhiteKing = &king
	board.BoardPieces[0][0] = &king
	pawn := newPawn(WHITE, Coordinate{1, 1})
	board.BoardPieces[1][1] = &pawn
	enemyQueen := newQueen(BLACK, Coordinate{2, 2})
	board.BoardPieces[2][2] = &enemyQueen
	validMoves := pawn.ValidMoves(&board)
	if len(validMoves) != 1 {
		t.Fatalf("Pawn should have one move")
	}
	_, ok := validMoves[Coordinate{2, 2}]
	if !ok {
		t.Fatalf("Pawn's only move should be to column 2, row 2")
	}
}
