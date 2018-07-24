// bishop_test
package chessgame

import (
	"testing"
)

func TestBishopCornerMoves(t *testing.T) {
	board := createEmptyBoard()
	bishop := newBishop(WHITE, Coordinate{0, 0})
	board.BoardPieces[0][0] = &bishop
	moves := bishop.ValidMoves(&board)
	if len(moves) != 7 {
		t.Fatalf("Expected 7 moves, got %d instead", len(moves))
	}
}

func TestBishopMovesTrappedByOwnPiece(t *testing.T) {
	board := createEmptyBoard()
	bishop := newBishop(WHITE, Coordinate{0, 0})
	board.BoardPieces[0][0] = &bishop
	king := newKing(WHITE, Coordinate{1, 1})
	board.BoardPieces[1][1] = &king
	moves := bishop.ValidMoves(&board)
	if len(moves) != 0 {
		t.Fatalf("Expected 0 moves, got %d instead", len(moves))
	}
}

func TestBishopMovesCapture(t *testing.T) {
	board := createEmptyBoard()
	bishop := newBishop(WHITE, Coordinate{0, 0})
	board.BoardPieces[0][0] = &bishop
	pawn := newPawn(BLACK, Coordinate{1, 1})
	board.BoardPieces[1][1] = &pawn
	moves := bishop.ValidMoves(&board)
	if len(moves) != 1 {
		t.Fatalf("Expected 1 moves, got %d instead", len(moves))
	}
}

func TestBishopMovesMiddleOfBoard(t *testing.T) {
	board := createEmptyBoard()
	bishop := newBishop(WHITE, Coordinate{3, 3})
	board.BoardPieces[3][3] = &bishop
	moves := bishop.ValidMoves(&board)
	if len(moves) != 13 {
		t.Fatalf("Expected 13 moves, got %d instead", len(moves))
	}
}
