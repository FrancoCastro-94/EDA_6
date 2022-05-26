package models

type Pawn struct {
	Board    *[17][17]string
	Row      int
	Col      int
	Side     string
	Opponent string
}

func (p *Pawn) Constructor(row, col int, side, opponent string, board *[17][17]string) {
	p.Board = board
	p.Row = row
	p.Col = col
	p.Side = side
	p.Opponent = opponent
}

// Return the position of the move and move prioriy
func (p *Pawn) GetMove() (to_row, to_col, priority int) {
	if p.Side == "N" {
		to_row, to_col, priority = p.getMoveNorthSide()
		return parsePosition(to_row), parsePosition(to_col), priority
	}
	to_row, to_col, priority = p.getMoveSouthSide()
	return parsePosition(to_row), parsePosition(to_col), priority
}

// Return the position of the move and move prioriy for N side
func (p *Pawn) getMoveNorthSide() (to_row, to_col, priority int) {
	priority = 0
	if p.canJumpOpponentToSouth() > priority {
		return p.Row + 4, p.Col, 2
	}
	if p.canMoveSouth() > priority {
		return p.Row + 2, p.Col, 1
	}
	moveSide := p.tryDodgeWall()
	if moveSide == "left" {
		return p.Row, p.Col - 2, 1
	}
	if moveSide == "right" {
		return p.Row, p.Col + 2, 1
	}
	if p.Row > 9 {
		if p.canJumpOpponentToLeft() > priority {
			return p.Row, p.Col - 4, 1
		}
		if p.canMoveLeft() > priority {
			return p.Row, p.Col - 2, 1
		}
		if p.canJumpOpponentToRigth() > priority {
			return p.Row, p.Col + 4, 0
		}
		if p.canMoveRight() > priority {
			return p.Row, p.Col + 2, 0
		}

	}
	if p.canJumpOpponentToRigth() > priority {
		return p.Row, p.Col + 4, 1
	}
	if p.canMoveRight() > priority {
		return p.Row, p.Col + 2, 1
	}
	if p.canJumpOpponentToLeft() > priority {
		return p.Row, p.Col - 4, 1
	}
	if p.canMoveLeft() > priority {
		return p.Row, p.Col - 2, 0
	}
	return 0, 0, -1
}

// Return the position of the move and move prioriy for S side
func (p *Pawn) getMoveSouthSide() (to_row, to_col, priority int) {
	priority = 0
	if p.canJumpOpponentToNorth() > priority {
		return p.Row - 4, p.Col, 2
	}
	if p.canMoveNorth() > priority {
		return p.Row - 2, p.Col, 1
	}
	moveSide := p.tryDodgeWall()
	if moveSide == "left" {
		return p.Row, p.Col - 2, 1
	}
	if moveSide == "right" {
		return p.Row, p.Col + 2, 1
	}
	if p.Row < 9 {
		if p.canJumpOpponentToRigth() > priority {
			return p.Row, p.Col + 4, 1
		}
		if p.canMoveRight() > priority {
			return p.Row, p.Col + 2, 1
		}
		if p.canJumpOpponentToLeft() > priority {
			return p.Row, p.Col - 4, 0
		}
		if p.canMoveLeft() > priority {
			return p.Row, p.Col - 2, 0
		}

	}
	if p.canJumpOpponentToLeft() > priority {
		return p.Row, p.Col - 4, 1
	}
	if p.canMoveLeft() > priority {
		return p.Row, p.Col - 2, 1
	}
	if p.canJumpOpponentToRigth() > priority {
		return p.Row, p.Col + 4, 0
	}
	if p.canMoveRight() > priority {
		return p.Row, p.Col + 2, 0
	}
	return 0, 0, -1
}

func (p *Pawn) canMoveRight() int {
	if p.isFreeRight(p.Row, p.Col, 2) {
		return 1
	}
	return 0
}

func (p *Pawn) canMoveLeft() int {
	if p.isFreeLeft(p.Row, p.Col, 2) {
		return 1
	}
	return 0
}

func (p *Pawn) canMoveSouth() int {
	if p.isFreeSouth(p.Row, p.Col, 2) {
		return 1
	}
	return -1
}
func (p *Pawn) canMoveNorth() int {
	if p.isFreeNorth(p.Row, p.Col, 2) {
		return 1
	}
	return -1
}

func (p *Pawn) canJumpOpponentToSouth() int {
	if p.isFreeSouth(p.Row+2, p.Col, 2) && p.Board[p.Row+1][p.Col] == " " && p.Board[p.Row+2][p.Col] == p.Opponent {
		return 2
	}
	return -1
}

func (p *Pawn) canJumpOpponentToNorth() int {
	if p.isFreeNorth(p.Row-2, p.Col, 2) && p.Board[p.Row-1][p.Col] == " " && p.Board[p.Row-2][p.Col] == p.Opponent {
		return 2
	}
	return -1
}

func (p *Pawn) canJumpOpponentToRigth() int {
	if p.isFreeRight(p.Row, p.Col+2, 2) && p.Board[p.Row][p.Col+1] == " " && p.Board[p.Row][p.Col+2] == p.Opponent {
		return 1
	}
	return -1
}

func (p *Pawn) canJumpOpponentToLeft() int {
	if p.isFreeLeft(p.Row, p.Col-2, 2) && p.Board[p.Row][p.Col-1] == " " && p.Board[p.Row][p.Col-2] == p.Opponent {
		return 1
	}
	return -1
}

// Return false if the row or column is out of board
func isInBoard(row, col int) bool {
	if row < 17 && col < 17 && row > -1 && col > -1 {
		return true
	}
	return false
}

// Returns true if is possible put wall and row and col of wall
func (p *Pawn) PossibleRightWall() (int, int, bool) {
	if p.isFreeNorth(p.Row-1, p.Col+1, 3) {
		return p.Row + 1, p.Col + 1, true
	}
	return -1, -1, false
}

// Returns true if is possible put wall and row and col of wall
func (p *Pawn) PutRightLeft() (int, int, bool) {
	if !isInBoard(p.Row+3, p.Col-1) {
		return -1, -1, false
	}
	if p.Board[p.Row+1][p.Col-1] == " " && p.Board[p.Row+2][p.Col-1] == " " && p.Board[p.Row-1][p.Col-1] == " " {
		return p.Row + 1, p.Col - 1, true
	}

	if p.Board[p.Row+1][p.Col-1] == " " && p.Board[p.Row+2][p.Col-1] == " " && p.Board[p.Row-1][p.Col-1] == " " {
		return p.Row + 1, p.Col - 1, true
	}
	return -1, -1, false
}

// Returns the position of front wall if is possible
func (p *Pawn) PossibleFrontWall() (int, int, bool) {
	if p.Side == "S" {
		if p.isFreeLeft(p.Row-1, p.Col+1, 3) {
			return parsePosition(p.Row), parsePosition(p.Col) - 1, true
		}
		return -1, -1, false
	} else {
		if isInBoard(p.Row+1, p.Col-2) && p.Board[p.Row+1][p.Col] == " " && p.Board[p.Row+1][p.Col-1] == " " && p.Board[p.Row+1][p.Col-2] == " " {
			return parsePosition(p.Row + 2), parsePosition(p.Col), true
		}
		if p.isFreeLeft(p.Row-1, p.Col+1, 3) {
			return parsePosition(p.Row), parsePosition(p.Col) - 1, true
		}
		return -1, -1, false
	}
}

// Checks the following places to the south are empty
func (p *Pawn) isFreeSouth(row, col, length int) bool {
	if isInBoard(row+length, col) {
		for i := 1; i <= length; i++ {
			if p.Board[row+i][col] != " " {
				return false
			}
		}
		return true
	}
	return false
}

// Checks the following places to the north are empty
func (p *Pawn) isFreeNorth(row, col, length int) bool {
	if isInBoard(row-length, col) {
		for i := 1; i <= length; i++ {
			if p.Board[row-i][col] != " " {
				return false
			}
		}
		return true
	}
	return false
}

// Checks the following places to the right are empty
func (p *Pawn) isFreeRight(row, col, length int) bool {
	if isInBoard(row, col+length) {
		for i := 1; i <= length; i++ {
			if p.Board[row][col+i] != " " {
				return false
			}
		}
		return true
	}
	return false
}

// Checks the following places to the left are empty
func (p *Pawn) isFreeLeft(row, col, length int) bool {
	if isInBoard(row, col-length) {
		for i := 1; i <= length; i++ {
			if p.Board[row][col-i] != " " {
				return false
			}
		}
		return true
	}
	return false
}

// Gets the number of row or column in the board 17x17 and return number for 9x9 board
func parsePosition(position int) int {
	return int(position / 2)
}

// Return the direction that need less changes to esquivar the wall
func (p *Pawn) tryDodgeWall() string {
	if p.Side == "S" {
		if isInBoard(p.Row-1, p.Col-3) && p.canMoveLeft() > 0 && p.Board[p.Row-1][p.Col-2] == " " {
			return "left"
		}
		if isInBoard(p.Row-1, p.Col+3) && p.canMoveRight() > 0 && p.Board[p.Row-1][p.Col+2] == " " {
			return "right"
		}
	} else {
		if isInBoard(p.Row+1, p.Col-3) && p.canMoveLeft() > 0 && p.Board[p.Row+1][p.Col-2] == " " {
			return "left"
		}
		if isInBoard(p.Row+1, p.Col+3) && p.canMoveRight() > 0 && p.Board[p.Row+1][p.Col+2] == " " {
			return "right"
		}
	}
	return ""
}
