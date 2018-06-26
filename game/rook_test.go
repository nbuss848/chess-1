// rook_test
package chessgame

import (
	"testing"
)

func TestRookMovesCorner(t *testing.T) {
	board := createEmptyBoard()
	rook := newRook(BLACK, Coordinate{0, 0})
	board.BoardPieces[0][0] = &rook
	moves := rook.validMoves(&board)
	if len(moves) != 14 {
		t.Fatalf("Expected 14 moves, got %d instead", len(moves))
	}
}

func TestRookMovesCapture(t *testing.T) {
	board := createEmptyBoard()
	blackRook := newRook(BLACK, Coordinate{0, 0})
	whiteQueen := newQueen(WHITE, Coordinate{0, 1})
	blackPawn := newPawn(BLACK, Coordinate{1, 0})
	board.BoardPieces[0][0] = &blackRook
	board.BoardPieces[1][0] = &blackPawn
	board.BoardPieces[0][1] = &whiteQueen
	moves := blackRook.validMoves(&board)
	if len(moves) != 1 {
		t.Fatalf("Expected 1 moves, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{0, 1}]
	if !ok {
		t.Fatalf("Row 0 column 1 should be only valid move")
	}
}

func TestRookMovesMiddleOfBoard(t *testing.T) {
	board := createEmptyBoard()
	rook := newRook(BLACK, Coordinate{4, 4})
	board.BoardPieces[4][4] = &rook
	moves := rook.validMoves(&board)
	if len(moves) != 14 {
		t.Fatalf("Expected 14 moves, got %d instead", len(moves))
	}
	pawn := newPawn(BLACK, Coordinate{2, 4})
	board.BoardPieces[2][4] = &pawn
	moves = rook.validMoves(&board)
	if len(moves) != 11 {
		t.Fatalf("Expected 11 moves, got %d instead", len(moves))
	}
}
