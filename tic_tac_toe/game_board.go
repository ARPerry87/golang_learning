package main

// This will represent the game board
type Board struct {
	Cells [3][3]string
}

// create a board with fully filled in cells with o's and x's

func filledBoard() Board {
	return Board{
		Cells: [3][3]string{
			{"X", "O", "X"},
			{"O", "X", "O"},
			{"O", "X", "X"},
		},
	}
}

func (b *Board) checkWin() string {
	// Check rows
	for i := range 3 {
		if b.Cells[i][0] == b.Cells[i][1] && b.Cells[i][1] == b.Cells[i][2] && b.Cells[i][0] != " " {
			return b.Cells[i][0]
		}
	}

	// Check columns
	for j := range 3 {
		if b.Cells[0][j] == b.Cells[1][j] && b.Cells[1][j] == b.Cells[2][j] && b.Cells[0][j] != " " {
			return b.Cells[0][j]
		}
	}

	// Check diagonals
	if b.Cells[0][0] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][2] && b.Cells[0][0] != " " {
		return b.Cells[0][0]
	}
	if b.Cells[0][2] == b.Cells[1][1] && b.Cells[1][1] == b.Cells[2][0] && b.Cells[0][2] != " " {
		return b.Cells[0][2]
	}

	return ""
}
