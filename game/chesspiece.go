// chesspiece
package chessgame

type Coordinate struct {
	Row    int
	Column int
}

type ChessPiece interface {
	ValidMoves(board *ChessBoard) map[Coordinate]bool
	updatePosition(coord Coordinate)
	GetPieceSide() Side
	GetPieceType() PieceType
	hasPieceMoved() bool
	GetCurrentCoordinates() Coordinate
	setHasMoved(hasMoved bool)
}

// Interface representation of piece that is threatening a King
type ThreateningPiece struct {
	coord     Coordinate
	pieceType PieceType
}

type Side int

const (
	WHITE Side = 1 + iota
	BLACK
)

type PieceType int

const (
	PAWN PieceType = 1 + iota
	BISHOP
	ROOK
	KNIGHT
	KING
	QUEEN
)
