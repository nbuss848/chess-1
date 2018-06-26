// chessboard_test
package chessgame

import (
	"testing"
)

func TestKingMovesMiddleOfBoard(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{4, 4})
	board.BoardPieces[4][4] = &king
	moves := king.validMoves(&board)
	if len(moves) != 8 {
		t.Fatalf("Expected 8 moves, got %d", len(moves))
	}
}

func TestKingMovesNoCheck(t *testing.T) {
	board := createEmptyBoard()
	king := newKing(WHITE, Coordinate{0, 4})
	board.BoardPieces[0][4] = &king
	moves := king.validMoves(&board)
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
	moves := king.validMoves(&board)
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
	moves := king.validMoves(&board)
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
	moves := king.validMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected no moves, got %d instead", len(moves))
	}
	queen.pieceSide = WHITE
	moves = king.validMoves(&board)
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
	moves := king.validMoves(&board)
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
	moves := king.validMoves(&board)
	firstCoord := Coordinate{0, 2}
	_, ok := moves[firstCoord]
	if !ok {
		t.Fatalf("Row 0 column 2 should be a valid move")
	}
	if len(moves) != 7 {
		t.Fatalf("Expected 7 moves, got %d instead", len(moves))
	}
}
