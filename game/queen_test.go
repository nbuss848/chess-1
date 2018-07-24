// queen_test
package chessgame

import (
	"testing"
)

func TestQueenMovesCorner(t *testing.T) {
	board := createEmptyBoard()
	queen := newQueen(WHITE, Coordinate{0, 0})
	board.BoardPieces[0][0] = &queen
	moves := queen.ValidMoves(&board)
	if len(moves) != 21 {
		t.Fatalf("Expected 21 moves, got %d instead", len(moves))
	}
}

func TestQueenMovesMiddle(t *testing.T) {
	board := createEmptyBoard()
	queen := newQueen(WHITE, Coordinate{4, 4})
	board.BoardPieces[4][4] = &queen
	moves := queen.ValidMoves(&board)
	if len(moves) != 27 {
		t.Fatalf("Expected 27 moves, got %d instead", len(moves))
	}
}

func TestQueenMovesTrappedByOwnPieces(t *testing.T) {
	board := createEmptyBoard()
	queen := newQueen(WHITE, Coordinate{0, 1})
	board.BoardPieces[0][1] = &queen
	pawn1 := newPawn(WHITE, Coordinate{0, 0})
	pawn2 := newPawn(WHITE, Coordinate{1, 0})
	pawn3 := newPawn(WHITE, Coordinate{0, 2})
	pawn4 := newPawn(WHITE, Coordinate{1, 2})
	board.BoardPieces[0][0] = &pawn1
	board.BoardPieces[1][0] = &pawn2
	board.BoardPieces[0][2] = &pawn3
	board.BoardPieces[1][2] = &pawn4
	moves := queen.ValidMoves(&board)
	if len(moves) != 7 {
		t.Fatalf("Expected 7 moves, got %d instead", len(moves))
	}
}

func TestQueenMovesCapture(t *testing.T) {
	board := createEmptyBoard()
	queen := newQueen(WHITE, Coordinate{0, 0})
	enemyKnight := newKnight(BLACK, Coordinate{1, 0})
	board.BoardPieces[0][0] = &queen
	board.BoardPieces[1][0] = &enemyKnight
	moves := queen.ValidMoves(&board)
	if len(moves) != 15 {
		t.Fatalf("Expected 15 moves, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{1, 0}]
	if !ok {
		t.Fatalf("Row 1, column 0 should be valid move")
	}
}

func TestQueenMovesOnlyDiagonal(t *testing.T) {
	board := CreateBoard()
	board.BoardPieces[6][3] = nil
	moves := board.BoardPieces[7][4].ValidMoves(&board)
	if len(moves) != 4 {
		t.Fatalf("Expected 4 moves, got %d instead", len(moves))
	}
}
