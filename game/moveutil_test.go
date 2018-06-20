// moveutil_test
package chessgame

import (
	"testing"
)

func TestGetNextCoordinate(t *testing.T) {
	coord := Coordinate{Row: 3, Column: 3}
	nextCoordinateRight := getNextStraightLineCoordinate(coord, 1, false)
	expectedRow := 3
	expectedCol := 4
	if expectedCol != nextCoordinateRight.Column {
		t.Fatalf("Expected %d but got %d", expectedCol, nextCoordinateRight.Column)
	}
	if expectedRow != nextCoordinateRight.Row {
		t.Fatalf("Expected %d but got %d", expectedRow, nextCoordinateRight.Row)
	}
}

func TestIsLegal(t *testing.T) {
	coord := Coordinate{Row: 3, Column: 3}
	isLegalMove := coord.isLegal()
	if !isLegalMove {
		t.Fatalf("Coordinate should be legal")
	}
	coord.Row = 7
	coord.Column = 0
	isLegalMove = coord.isLegal()
	if !isLegalMove {
		t.Fatalf("Coordinate should be legal")
	}
	coord.Row = 8
	isLegalMove = coord.isLegal()
	if isLegalMove {
		t.Fatalf("Coordinate should be illegal")
	}
	coord.Row = 0
	isLegalMove = coord.isLegal()
	if !isLegalMove {
		t.Fatalf("Coordinate should be legal")
	}
	coord.Column = -1
	isLegalMove = coord.isLegal()
	if isLegalMove {
		t.Fatalf("Coordinate should be illegal")
	}
}
