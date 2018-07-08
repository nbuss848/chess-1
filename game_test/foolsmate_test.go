// foolsmate_test
package game_test

import (
	"chess/game"

	"testing"
)

type move struct {
	fromCoord chessgame.Coordinate
	toCoord   chessgame.Coordinate
}

type FakePlayer struct {
	coords    []move
	moveIndex int
	side      chessgame.Side
}

func (player *FakePlayer) MakeMove(board chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	player.moveIndex += 1
	return player.coords[player.moveIndex-1].fromCoord, player.coords[player.moveIndex-1].toCoord
}

func (player *FakePlayer) GetSide() chessgame.Side {
	return player.side
}

func (player *FakePlayer) PromotePawn() chessgame.PieceType {
	return chessgame.QUEEN
}

func TestGameSetup(t *testing.T) {
	moves := []move{move{chessgame.Coordinate{1, 1}, chessgame.Coordinate{2, 2}}}
	player1 := FakePlayer{moves, 0, chessgame.WHITE}
	player2 := FakePlayer{moves, 0, chessgame.BLACK}
	game := chessgame.NewChessGame(&player1, &player2)
	if game.WhitePlayer.GetSide() != chessgame.WHITE {
		t.Fatalf("game setup failed")
	}
}

func TestGameFoolsMate(t *testing.T) {
	whiteMove1 := move{chessgame.Coordinate{1, 2}, chessgame.Coordinate{2, 2}}
	whiteMove2 := move{chessgame.Coordinate{1, 1}, chessgame.Coordinate{3, 1}}
	blackMove1 := move{chessgame.Coordinate{6, 3}, chessgame.Coordinate{4, 3}}
	blackMove2 := move{chessgame.Coordinate{7, 4}, chessgame.Coordinate{3, 0}}
	whiteMoves := []move{whiteMove1, whiteMove2}
	blackMoves := []move{blackMove1, blackMove2}
	player1 := FakePlayer{whiteMoves, 0, chessgame.WHITE}
	player2 := FakePlayer{blackMoves, 0, chessgame.BLACK}
	game := chessgame.NewChessGame(&player1, &player2)
	if game.PlayGame() != chessgame.BLACKVICTORY {
		t.Fatalf("game should result in black checkmating white")
	}
}

func TestGameFoolsMateWithBadMoves(t *testing.T) {
	whiteMove1 := move{chessgame.Coordinate{0, 4}, chessgame.Coordinate{7, 4}}
	whiteMove2 := move{chessgame.Coordinate{0, 0}, chessgame.Coordinate{1, 0}}
	whiteMove3 := move{chessgame.Coordinate{1, 2}, chessgame.Coordinate{2, 2}}
	whiteMove4 := move{chessgame.Coordinate{1, 1}, chessgame.Coordinate{3, 1}}
	blackMove1 := move{chessgame.Coordinate{0, 3}, chessgame.Coordinate{0, 4}}
	blackMove2 := move{chessgame.Coordinate{6, 3}, chessgame.Coordinate{4, 3}}
	blackMove3 := move{chessgame.Coordinate{7, 4}, chessgame.Coordinate{3, 0}}
	whiteMoves := []move{whiteMove1, whiteMove2, whiteMove3, whiteMove4}
	blackMoves := []move{blackMove1, blackMove2, blackMove3}
	player1 := FakePlayer{whiteMoves, 0, chessgame.WHITE}
	player2 := FakePlayer{blackMoves, 0, chessgame.BLACK}
	game := chessgame.NewChessGame(&player1, &player2)
	if game.PlayGame() != chessgame.BLACKVICTORY {
		t.Fatalf("game should result in black checkmating white")
	}
}
