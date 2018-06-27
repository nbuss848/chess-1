// knight_test
package chessgame

import (
	"testing"
)

func TestKnightMovesEdge(t *testing.T) {
	board := createEmptyBoard()
	knight := newKnight(BLACK, Coordinate{0, 4})
	board.BoardPieces[0][4] = &knight
	moves := knight.validMoves(&board)
	if len(moves) != 4 {
		t.Fatalf("Expected 4 moves, got %d instead", len(moves))
	}
}

func TestKnightMovesMiddle(t *testing.T) {
	board := createEmptyBoard()
	knight := newKnight(BLACK, Coordinate{4, 4})
	board.BoardPieces[4][4] = &knight
	moves := knight.validMoves(&board)
	if len(moves) != 8 {
		t.Fatalf("Expected 8 moves, got %d instead", len(moves))
	}
}

func TestKnightMovesMiddleCapture(t *testing.T) {
	board := createEmptyBoard()
	knight := newKnight(BLACK, Coordinate{4, 4})
	board.BoardPieces[4][4] = &knight
	bishopToCapture := newBishop(WHITE, Coordinate{2, 3})
	board.BoardPieces[2][3] = &bishopToCapture
	moves := knight.validMoves(&board)
	if len(moves) != 8 {
		t.Fatalf("Expected 8 moves, got %d instead", len(moves))
	}
	expectedCoord := Coordinate{2, 3}
	_, ok := moves[expectedCoord]
	if !ok {
		t.Fatalf("Row 2 column 3 should be in valid moves")
	}
}

func TestKnightMovesBlockedMove(t *testing.T) {
	board := createEmptyBoard()
	knight := newKnight(BLACK, Coordinate{4, 4})
	board.BoardPieces[4][4] = &knight
	friendlyBishop := newBishop(BLACK, Coordinate{2, 3})
	board.BoardPieces[2][3] = &friendlyBishop
	moves := knight.validMoves(&board)
	if len(moves) != 7 {
		t.Fatalf("Expected 7 moves, got %d instead", len(moves))
	}
	expectedCoord := Coordinate{2, 3}
	_, ok := moves[expectedCoord]
	if ok {
		t.Fatalf("Row 2 column 3 should not be in valid moves")
	}
}
