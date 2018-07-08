// scholarsmate_test
package game_test

import (
	"chess/game"

	"testing"
)

func TestGameScholarsMate(t *testing.T) {
	whiteMove1 := move{chessgame.Coordinate{1, 3}, chessgame.Coordinate{3, 3}}
	blackMove1 := move{chessgame.Coordinate{6, 5}, chessgame.Coordinate{4, 5}}
	whiteMove2 := move{chessgame.Coordinate{0, 2}, chessgame.Coordinate{3, 5}}
	blackMove2 := move{chessgame.Coordinate{7, 6}, chessgame.Coordinate{5, 5}}
	whiteMove3 := move{chessgame.Coordinate{0, 4}, chessgame.Coordinate{4, 0}}
	blackMove3 := move{chessgame.Coordinate{7, 1}, chessgame.Coordinate{5, 2}}
	whiteMove4 := move{chessgame.Coordinate{4, 0}, chessgame.Coordinate{6, 2}}
	whiteMoves := []move{whiteMove1, whiteMove2, whiteMove3, whiteMove4}
	blackMoves := []move{blackMove1, blackMove2, blackMove3}
	player1 := FakePlayer{whiteMoves, 0, chessgame.WHITE}
	player2 := FakePlayer{blackMoves, 0, chessgame.BLACK}
	game := chessgame.NewChessGame(&player1, &player2)
	if game.PlayGame() != chessgame.WHITEVICTORY {
		t.Fatalf("game should result in white checkmating black")
	}
}

func TestGameScholarsMateWithBadMoves(t *testing.T) {
	whiteMove1 := move{chessgame.Coordinate{1, 3}, chessgame.Coordinate{3, 3}}
	blackMove1 := move{chessgame.Coordinate{6, 3}, chessgame.Coordinate{4, 3}}
	whiteMove2 := move{chessgame.Coordinate{3, 3}, chessgame.Coordinate{4, 3}}
	whiteMove3 := move{chessgame.Coordinate{0, 2}, chessgame.Coordinate{3, 5}}
	blackMove2 := move{chessgame.Coordinate{7, 6}, chessgame.Coordinate{5, 5}}
	whiteMove4 := move{chessgame.Coordinate{0, 4}, chessgame.Coordinate{4, 0}}
	blackMove3 := move{chessgame.Coordinate{6, 2}, chessgame.Coordinate{5, 2}}
	blackMove4 := move{chessgame.Coordinate{7, 1}, chessgame.Coordinate{5, 2}}
	whiteMove5 := move{chessgame.Coordinate{4, 0}, chessgame.Coordinate{6, 2}}
	whiteMoves := []move{whiteMove1, whiteMove2, whiteMove3, whiteMove4, whiteMove5}
	blackMoves := []move{blackMove1, blackMove2, blackMove3, blackMove4}
	player1 := FakePlayer{whiteMoves, 0, chessgame.WHITE}
	player2 := FakePlayer{blackMoves, 0, chessgame.BLACK}
	game := chessgame.NewChessGame(&player1, &player2)
	if game.PlayGame() != chessgame.WHITEVICTORY {
		t.Fatalf("game should result in white checkmating black")
	}
}
