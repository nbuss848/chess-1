// movesimulator
package chessengine

import (
	"chess/game"
)

func selectBestMove(side chessgame.Side, board *chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	maxScore := 0
	var bestFromCoord chessgame.Coordinate
	var bestToCoord chessgame.Coordinate
	for from, allToCoords := range validMoves {
		for to, _ := range allToCoords {
			board.UpdateBoard(from, to)
			score := 0
			if board.BoardPieces[to.Row][to.Column] != nil && board.GetPieceSide(to) != side {
				score += pieceValue(board.BoardPieces[to.Row][to.Column]) * 20
			}
			score += evaluateBoard(board, side)
			if score >= maxScore {
				maxScore = score
				bestFromCoord = from
				bestToCoord = to
			}
			board.UpdateBoard(to, from)
		}
	}
	return bestFromCoord, bestToCoord
}
