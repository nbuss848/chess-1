// chessboard_test
package chessgame

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	board := CreateBoard()
	if len(board.BoardPieces) != 8 {
		t.Fatalf("Expected length of 8, but got %d", len(board.BoardPieces))
	}
	if len(board.BoardPieces[0]) != 8 {
		t.Fatalf("Expected length of 8, but got %d", len(board.BoardPieces[0]))
	}
}

func TestCreateBoardBaseRows(t *testing.T) {
	board := CreateBoard()
	for i := 0; i < 8; i++ {
		shouldBeWhitePiece := board.BoardPieces[0][i].getPieceSide()
		if shouldBeWhitePiece != WHITE {
			t.Fatalf("Piece should be white")
		}
		shouldBeBlackPiece := board.BoardPieces[7][1].getPieceSide()
		if shouldBeBlackPiece != BLACK {
			t.Fatalf("Piece should be black")
		}
	}
	pieceTypeWhite := board.BoardPieces[0][0].getPieceType()
	pieceTypeBlack := board.BoardPieces[7][0].getPieceType()
	if pieceTypeWhite != ROOK && pieceTypeBlack != ROOK {
		t.Fatalf("Expected ROOK")
	}
	pieceTypeWhite = board.BoardPieces[0][1].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][1].getPieceType()
	if pieceTypeWhite != KNIGHT || pieceTypeBlack != KNIGHT {
		t.Fatalf("Expected KNIGHT")
	}
	pieceTypeWhite = board.BoardPieces[0][2].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][2].getPieceType()
	if pieceTypeWhite != BISHOP || pieceTypeBlack != BISHOP {
		t.Fatalf("Expected BISHOP")
	}
	pieceTypeWhite = board.BoardPieces[0][3].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][3].getPieceType()
	if pieceTypeWhite != QUEEN || pieceTypeBlack != QUEEN {
		t.Fatalf("Expected QUEEN")
	}
	pieceTypeWhite = board.BoardPieces[0][4].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][4].getPieceType()
	if pieceTypeWhite != KING || pieceTypeBlack != KING {
		t.Fatalf("Expected KING")
	}
	pieceTypeWhite = board.BoardPieces[0][5].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][5].getPieceType()
	if pieceTypeWhite != BISHOP || pieceTypeBlack != BISHOP {
		t.Fatalf("Expected BISHOP")
	}
	pieceTypeWhite = board.BoardPieces[0][6].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][6].getPieceType()
	if pieceTypeWhite != KNIGHT || pieceTypeBlack != KNIGHT {
		t.Fatalf("Expected KNIGHT")
	}
	pieceTypeWhite = board.BoardPieces[0][7].getPieceType()
	pieceTypeBlack = board.BoardPieces[7][7].getPieceType()
	if pieceTypeWhite != ROOK || pieceTypeBlack != ROOK {
		t.Fatalf("Expected ROOK")
	}
}

func TestCreateBoardPawns(t *testing.T) {
	board := CreateBoard()
	for i := 0; i < 8; i++ {
		shouldBeWhitePawn := board.BoardPieces[1][i]
		if shouldBeWhitePawn.getPieceSide() != WHITE {
			t.Fatalf("Side should be white")
		}
		if shouldBeWhitePawn.getPieceType() != PAWN {
			t.Fatalf("Piece type should be pawn")
		}
		shouldBeBlackPawn := board.BoardPieces[6][i]
		if shouldBeBlackPawn.getPieceSide() != BLACK {
			t.Fatalf("Side should be black")
		}
		if shouldBeBlackPawn.getPieceType() != PAWN {
			t.Fatalf("Piece type should be pawn")
		}
	}
}

func TestCreateBoardEmptySpaces(t *testing.T) {
	board := CreateBoard()
	for i := 2; i < 6; i++ {
		for j := 0; j < 8; j++ {
			if board.BoardPieces[i][j] != nil {
				t.Fatalf("Space should be nil")
			}
		}
	}
}
