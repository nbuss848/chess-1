// commandlineai
package main

import (
	"chess/engine"
	"chess/game"
	"fmt"
	"time"
)

type CommandLineAI struct {
	side chessgame.Side
}

func (ai CommandLineAI) MakeMove(boardClone chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	side := chessgame.WHITE
	if ai.GetSide() == chessgame.WHITE {
		side = chessgame.BLACK
	}
	PrintBoard(boardClone, side)
	duration := time.Second
	time.Sleep(duration)
	fmt.Println("\n")
	return chessengine.MakeEvaluatedMove(ai.GetSide(), &boardClone, validMoves)
}

func (ai CommandLineAI) PromotePawn() chessgame.PieceType {
	return chessgame.QUEEN
}

func (ai CommandLineAI) GetSide() chessgame.Side {
	return ai.side
}

func NewCommandLineAI(side chessgame.Side) CommandLineAI {
	return CommandLineAI{side}
}
