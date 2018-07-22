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
		shouldBeWhitePiece := board.BoardPieces[0][i].GetPieceSide()
		if shouldBeWhitePiece != WHITE {
			t.Fatalf("Piece should be white")
		}
		shouldBeBlackPiece := board.BoardPieces[7][1].GetPieceSide()
		if shouldBeBlackPiece != BLACK {
			t.Fatalf("Piece should be black")
		}
	}
	pieceTypeWhite := board.BoardPieces[0][0].GetPieceType()
	pieceTypeBlack := board.BoardPieces[7][0].GetPieceType()
	if pieceTypeWhite != ROOK && pieceTypeBlack != ROOK {
		t.Fatalf("Expected ROOK")
	}
	pieceTypeWhite = board.BoardPieces[0][1].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][1].GetPieceType()
	if pieceTypeWhite != KNIGHT || pieceTypeBlack != KNIGHT {
		t.Fatalf("Expected KNIGHT")
	}
	pieceTypeWhite = board.BoardPieces[0][2].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][2].GetPieceType()
	if pieceTypeWhite != BISHOP || pieceTypeBlack != BISHOP {
		t.Fatalf("Expected BISHOP")
	}
	pieceTypeWhite = board.BoardPieces[0][3].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][3].GetPieceType()
	if pieceTypeWhite != KING || pieceTypeBlack != KING {
		t.Fatalf("Expected KING")
	}
	pieceTypeWhite = board.BoardPieces[0][4].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][4].GetPieceType()
	if pieceTypeWhite != QUEEN || pieceTypeBlack != QUEEN {
		t.Fatalf("Expected QUEEN")
	}
	pieceTypeWhite = board.BoardPieces[0][5].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][5].GetPieceType()
	if pieceTypeWhite != BISHOP || pieceTypeBlack != BISHOP {
		t.Fatalf("Expected BISHOP")
	}
	pieceTypeWhite = board.BoardPieces[0][6].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][6].GetPieceType()
	if pieceTypeWhite != KNIGHT || pieceTypeBlack != KNIGHT {
		t.Fatalf("Expected KNIGHT")
	}
	pieceTypeWhite = board.BoardPieces[0][7].GetPieceType()
	pieceTypeBlack = board.BoardPieces[7][7].GetPieceType()
	if pieceTypeWhite != ROOK || pieceTypeBlack != ROOK {
		t.Fatalf("Expected ROOK")
	}
}

func TestCreateBoardPawns(t *testing.T) {
	board := CreateBoard()
	for i := 0; i < 8; i++ {
		shouldBeWhitePawn := board.BoardPieces[1][i]
		if shouldBeWhitePawn.GetPieceSide() != WHITE {
			t.Fatalf("Side should be white")
		}
		if shouldBeWhitePawn.GetPieceType() != PAWN {
			t.Fatalf("Piece type should be pawn")
		}
		shouldBeBlackPawn := board.BoardPieces[6][i]
		if shouldBeBlackPawn.GetPieceSide() != BLACK {
			t.Fatalf("Side should be black")
		}
		if shouldBeBlackPawn.GetPieceType() != PAWN {
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

func TestDeepCopyBoard(t *testing.T) {
	board := CreateBoard()
	clonedBoard := deepCopyBoard(&board)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			coord := Coordinate{row, col}
			if board.isSpaceOccupied(coord) != clonedBoard.isSpaceOccupied(coord) {
				t.Fatalf("Row: %d Column: %d does not match", row, col)
			}
			if !board.isSpaceOccupied(coord) && !clonedBoard.isSpaceOccupied(coord) {
				continue
			}
			if board.GetPieceSide(coord) != clonedBoard.GetPieceSide(coord) {
				t.Fatalf("Row: %d Column: %d for boards do not have matching sides", row, col)
			}
			if board.GetPieceType(coord) != clonedBoard.GetPieceType(coord) {
				t.Fatalf("Row: %d Column: %d for boards do not have matching piece types", row, col)
			}
		}
	}
}
