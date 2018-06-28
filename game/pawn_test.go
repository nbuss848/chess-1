// pawn_test.go
package chessgame

import (
	"testing"
)

func TestBlackPawnInitialMoves(t *testing.T) {
	board := createEmptyBoard()
	pawn := newPawn(BLACK, Coordinate{6, 4})
	board.BoardPieces[6][4] = &pawn
	moves := pawn.validMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 2 moves, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{5, 4}]
	if !ok {
		t.Fatalf("Row 5 column 4 should be valid move")
	}
	_, ok = moves[Coordinate{4, 4}]
	if !ok {
		t.Fatalf("Row 4 column 4 should be validm move")
	}
}

func TestWhitePawnInitialMoves(t *testing.T) {
	board := createEmptyBoard()
	pawn := newPawn(WHITE, Coordinate{1, 4})
	board.BoardPieces[1][4] = &pawn
	moves := pawn.validMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 2 moves, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{2, 4}]
	if !ok {
		t.Fatalf("Row 2 column 4 should be valid move")
	}
	_, ok = moves[Coordinate{3, 4}]
	if !ok {
		t.Fatalf("Row 3 column 4 should be validm move")
	}
}

func TestPawnMovesAfterInitialMove(t *testing.T) {
	board := createEmptyBoard()
	pawn := newPawn(WHITE, Coordinate{2, 4})
	pawn.hasMoved = true
	board.BoardPieces[2][4] = &pawn
	moves := pawn.validMoves(&board)
	if len(moves) != 1 {
		t.Fatalf("Expected 1 moves, got %d instead", len(moves))
	}
}

func TestPawnMovesCapture(t *testing.T) {
	board := createEmptyBoard()
	pawn := newPawn(WHITE, Coordinate{2, 4})
	pawn.hasMoved = true
	board.BoardPieces[2][4] = &pawn
	enemyPawn1 := newPawn(BLACK, Coordinate{3, 3})
	enemyPawn2 := newPawn(BLACK, Coordinate{3, 5})
	board.BoardPieces[3][3] = &enemyPawn1
	board.BoardPieces[3][5] = &enemyPawn2
	moves := pawn.validMoves(&board)
	if len(moves) != 3 {
		t.Fatalf("Expected 3 moves, got %d instead", len(moves))
	}
	enemyPawn1.pieceSide = WHITE
	moves = pawn.validMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 2 moves, got %d instead", len(moves))
	}
}

func TestPawnMovesEnPassant(t *testing.T) {
	board := createEmptyBoard()
	pawn := newPawn(WHITE, Coordinate{4, 4})
	pawn.hasMoved = true
	board.BoardPieces[4][4] = &pawn
	enemyPawn := newPawn(BLACK, Coordinate{4, 5})
	board.BoardPieces[4][5] = &enemyPawn
	board.MoveLog = append(board.MoveLog, Move{Coordinate{6, 5}, Coordinate{4, 5}, false, &enemyPawn})
	moves := pawn.validMoves(&board)
	if len(moves) != 2 {
		t.Fatalf("Expected 2 moves, got %d instead", len(moves))
	}
	_, ok := moves[Coordinate{5, 5}]
	if !ok {
		t.Fatalf("Row 3 column 5 should be valid move (en passant)")
	}
}
