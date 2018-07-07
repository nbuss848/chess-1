// chessboard
package chessgame

// Struct representing chess board
type ChessBoard struct {
	BoardPieces [][]ChessPiece
	MoveLog     []Move
	WhiteKing   *King
	BlackKing   *King
}

// Represents a move in a chess game
type Move struct {
	fromCoordinate Coordinate
	toCoordinate   Coordinate
	isCapture      bool
	pieceType      PieceType
}

// Calls on player to make turn, and updates board accordingly. Returns TurnOutcome, which can be a White Checkmate, Black Checkmate, Stalemate, or Continue
func (board *ChessBoard) TakeTurn(player ChessPlayer) TurnOutcome {
	inCheck := false
	if player.GetSide() == WHITE {
		board.WhiteKing.updateKingStatus(board)
		inCheck = board.WhiteKing.inCheck
	} else {
		board.BlackKing.updateKingStatus(board)
		inCheck = board.BlackKing.inCheck
	}
	side := player.GetSide()
	allValidMoves := board.getAllValidMovesForSide(side)
	if inCheck && len(allValidMoves) == 0 {
		if player.GetSide() == WHITE {
			return BLACKCHECKMATE
		} else {
			return WHITECHECKMATE
		}
	}
	if len(allValidMoves) == 0 {
		return STALEMATE
	}
	clonedBoard := deepCopyBoard(board)
	fromCoord, toCoord := player.MakeMove(clonedBoard, allValidMoves)
	validMoves, fromOk := allValidMoves[fromCoord]
	toOk := false
	if fromOk {
		_, toOk = validMoves[toCoord]
	}
	for !fromOk || !toOk {
		fromCoord, toCoord = player.MakeMove(clonedBoard, allValidMoves)
		validMoves, fromOk = allValidMoves[fromCoord]
		if fromOk {
			_, toOk = validMoves[toCoord]
		}
	}
	promotePawn := board.updateBoard(fromCoord, toCoord)
	// Update side's king to not be in check, since move validation wouldn't allow king to still be in check if it were before
	if side == WHITE {
		board.WhiteKing.inCheck = false
		board.WhiteKing.threateningPieces = nil
	} else {
		board.BlackKing.inCheck = false
		board.BlackKing.threateningPieces = nil
	}
	if !promotePawn {
		return CONTINUE
	}
	promotePiece := player.PromotePawn()
	for promotePiece == KING || promotePiece == PAWN {
		promotePiece = player.PromotePawn()
	}
	board.promotePawn(toCoord, promotePiece, side)
	return CONTINUE
}

// Returns all valid moves for given side, as a dictionary with coordinates as keys and maps of coordinates and booleans as the values
func (board *ChessBoard) getAllValidMovesForSide(pieceSide Side) map[Coordinate]map[Coordinate]bool {
	moves := make(map[Coordinate]map[Coordinate]bool)
	for row := 0; row < len(board.BoardPieces); row++ {
		for col := 0; col < len(board.BoardPieces[row]); col++ {
			currentCoord := Coordinate{row, col}
			if board.isSpaceOccupied(currentCoord) && board.getPieceSide(currentCoord) == pieceSide {
				pieceMoves := board.BoardPieces[row][col].validMoves(board)
				if len(pieceMoves) != 0 {
					moves[currentCoord] = pieceMoves
				}
			}
		}
	}
	return moves
}

// Given valid move (in form of from Coordinate and to Coordinate) updates board and returns true if pawn needs to be promoted, false otherwise
func (board *ChessBoard) updateBoard(fromCoord Coordinate, toCoord Coordinate) bool {
	piece := board.getPieceType(fromCoord)
	isCapture := false
	if board.isSpaceOccupied(toCoord) {
		isCapture = true
	}
	enPassant := false
	if piece == PAWN && toCoord.Column != fromCoord.Column && !board.isSpaceOccupied(toCoord) {
		enPassant = true
		isCapture = true
	}
	board.BoardPieces[toCoord.Row][toCoord.Column] = board.BoardPieces[fromCoord.Row][fromCoord.Column]
	board.BoardPieces[fromCoord.Row][fromCoord.Column] = nil
	board.BoardPieces[toCoord.Row][toCoord.Column].updatePosition(toCoord)
	if piece == KING && fromCoord.Column-toCoord.Column > 1 {
		newRookCoord := Coordinate{toCoord.Row, toCoord.Column + 1}
		board.BoardPieces[newRookCoord.Row][newRookCoord.Column] = board.BoardPieces[newRookCoord.Row][0]
		board.BoardPieces[newRookCoord.Row][0] = nil
		board.BoardPieces[newRookCoord.Row][newRookCoord.Column].updatePosition(newRookCoord)
	}
	if piece == KING && fromCoord.Column-toCoord.Column < 1 {
		newRookCoord := Coordinate{toCoord.Row, toCoord.Column - 1}
		board.BoardPieces[newRookCoord.Row][newRookCoord.Column] = board.BoardPieces[newRookCoord.Row][7]
		board.BoardPieces[newRookCoord.Row][0] = nil
		board.BoardPieces[newRookCoord.Row][newRookCoord.Column].updatePosition(newRookCoord)
	}
	if enPassant {
		board.BoardPieces[fromCoord.Row][toCoord.Column] = nil
	}
	if piece == PAWN && (toCoord.Row == 0 || toCoord.Row == 7) {
		return true
	}
	board.MoveLog = append(board.MoveLog, Move{fromCoord, toCoord, isCapture, board.getPieceType(toCoord)})
	return false
}

func (board *ChessBoard) promotePawn(coord Coordinate, pType PieceType, side Side) {
	var newPiece ChessPiece
	if pType == KNIGHT {
		knight := newKnight(side, coord)
		newPiece = &knight
	} else if pType == BISHOP {
		bishop := newBishop(side, coord)
		newPiece = &bishop
	} else if pType == ROOK {
		rook := newRook(side, coord)
		newPiece = &rook
	} else {
		queen := newQueen(side, coord)
		newPiece = &queen
	}
	board.BoardPieces[coord.Row][coord.Column] = newPiece
}

// Creates new ChessBoard with all pieces in place for beginning of game
func CreateBoard() ChessBoard {
	MoveLog := []Move{}
	whiteBaseRow := createBaseRow(0, WHITE)
	whitePawns := createPawnRow(1, WHITE)
	emptyRow1 := createEmptyRow()
	emptyRow2 := createEmptyRow()
	emptyRow3 := createEmptyRow()
	emptyRow4 := createEmptyRow()
	blackPawns := createPawnRow(6, BLACK)
	blackBaseRow := createBaseRow(7, BLACK)
	board := [][]ChessPiece{whiteBaseRow, whitePawns, emptyRow1, emptyRow2, emptyRow3, emptyRow4, blackPawns, blackBaseRow}
	whiteKing := board[0][3].(*King)
	blackKing := board[7][3].(*King)
	return ChessBoard{board, MoveLog, whiteKing, blackKing}
}

// Creates base row for given side, with rooks, knights, bishops, king and queen
func createBaseRow(row int, pieceSide Side) []ChessPiece {
	leftRook := newRook(pieceSide, Coordinate{Row: row, Column: 0})
	leftKnight := newKnight(pieceSide, Coordinate{Row: row, Column: 1})
	leftbishop := newBishop(pieceSide, Coordinate{Row: row, Column: 2})
	king := newKing(pieceSide, Coordinate{Row: row, Column: 3})
	queen := newQueen(pieceSide, Coordinate{Row: row, Column: 4})
	rightBishop := newBishop(pieceSide, Coordinate{Row: row, Column: 5})
	rightKnight := newKnight(pieceSide, Coordinate{Row: row, Column: 6})
	rightRook := newRook(pieceSide, Coordinate{Row: row, Column: 7})
	return []ChessPiece{&leftRook, &leftKnight, &leftbishop, &king, &queen, &rightBishop, &rightKnight, &rightRook}
}

// Creates a deep copy of the given ChessBoard
func deepCopyBoard(board *ChessBoard) ChessBoard {
	var whiteClonedKing *King
	var blackClonedKing *King
	clonedBoardPieces := make([][]ChessPiece, 8, 8)
	for i := 0; i < len(clonedBoardPieces); i++ {
		clonedBoardPieces[i] = make([]ChessPiece, 8, 8)
	}
	for row := 0; row < len(board.BoardPieces); row++ {
		for col := 0; col < len(board.BoardPieces[row]); col++ {
			piece := board.BoardPieces[row][col]
			var clonedPiece ChessPiece
			if piece == nil {
				clonedBoardPieces[row][col] = nil
				continue
			}
			pieceType := piece.getPieceType()
			pieceSide := piece.getPieceSide()
			currentCoord := piece.getCurrentCoordinates()
			if pieceType == PAWN {
				pawn := newPawn(pieceSide, currentCoord)
				clonedPiece = &pawn
			} else if pieceType == ROOK {
				rook := newRook(pieceSide, currentCoord)
				clonedPiece = &rook
			} else if pieceType == KNIGHT {
				knight := newKnight(pieceSide, currentCoord)
				clonedPiece = &knight
			} else if pieceType == BISHOP {
				bishop := newBishop(pieceSide, currentCoord)
				clonedPiece = &bishop
			} else if pieceType == QUEEN {
				queen := newQueen(pieceSide, currentCoord)
				clonedPiece = &queen
			} else {
				kingStruct := piece.(*King)
				king := newKing(pieceSide, currentCoord)
				clonedPiece = &king
				clonedPiece.(*King).inCheck = kingStruct.inCheck
				for _, tPiece := range kingStruct.threateningPieces {
					clonedPiece.(*King).threateningPieces = append(clonedPiece.(*King).threateningPieces, tPiece)
				}
				if pieceSide == BLACK {
					blackClonedKing = clonedPiece.(*King)
				} else {
					whiteClonedKing = clonedPiece.(*King)
				}
			}
			if piece.hasPieceMoved() {
				clonedPiece.setHasMoved(true)
			}
			clonedBoardPieces[row][col] = clonedPiece
		}
	}
	return ChessBoard{BoardPieces: clonedBoardPieces, WhiteKing: whiteClonedKing, BlackKing: blackClonedKing}
}

// Creates row of pawns for ChessBoard, all of given side
func createPawnRow(row int, pieceSide Side) []ChessPiece {
	var pawns []ChessPiece
	for i := 0; i < 8; i++ {
		pawn := newPawn(pieceSide, Coordinate{Row: row, Column: i})
		pawns = append(pawns, &pawn)
	}
	return pawns
}

// Creates empty row for ChessBoard
func createEmptyRow() []ChessPiece {
	return []ChessPiece{nil, nil, nil, nil, nil, nil, nil, nil}
}

// Creates empty ChessBoard - for testing purposes
func createEmptyBoard() ChessBoard {
	emptyBoard := [][]ChessPiece{}
	for i := 0; i < 8; i++ {
		emptyBoard = append(emptyBoard, createEmptyRow())
	}
	MoveLog := []Move{}
	// Assign kings fake coordinates for testing purposes
	whiteKing := newKing(WHITE, Coordinate{-1, -1})
	blackKing := newKing(BLACK, Coordinate{8, 8})
	return ChessBoard{emptyBoard, MoveLog, &whiteKing, &blackKing}
}

// Checks whether given space on board is occupied
func (board *ChessBoard) isSpaceOccupied(coord Coordinate) bool {
	if board.BoardPieces[coord.Row][coord.Column] != nil {
		return true
	}
	return false
}

// Gets type of piece occupying given coordinate
func (board *ChessBoard) getPieceType(coord Coordinate) PieceType {
	return board.BoardPieces[coord.Row][coord.Column].getPieceType()
}

// Gets side of piece occupying given coordinate
func (board *ChessBoard) getPieceSide(coord Coordinate) Side {
	return board.BoardPieces[coord.Row][coord.Column].getPieceSide()
}

// Gets previous move on board
func (board *ChessBoard) getPreviousMove() (Move, bool) {
	if len(board.MoveLog) == 0 {
		return Move{}, false
	}
	return board.MoveLog[len(board.MoveLog)-1], true
}

// Checks whether piece occupying given coordinate has moved
func (board *ChessBoard) hasPieceMoved(coord Coordinate) bool {
	return board.BoardPieces[coord.Row][coord.Column].hasPieceMoved()
}

// Returns Coordinate of king for given side
func (board *ChessBoard) getKingCoordinate(pieceSide Side) Coordinate {
	if pieceSide == WHITE {
		return board.WhiteKing.currentCoordinate
	}
	return board.BlackKing.currentCoordinate
}
