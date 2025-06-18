package main

const (
	PlayerX = "X"
	PlayerO = "O"
	Empty   = " "
)

// This will represent the game board
type Board struct {
	Cells [3][3]string
}

// create a board with fully filled in cells with o's and x's
func filledBoard() Board {
	return Board{
		Cells: [3][3]string{
			{PlayerX, PlayerO, PlayerX},
			{PlayerO, PlayerX, PlayerO},
			{PlayerO, PlayerX, PlayerX},
		},
	}
}

func (b *Board) checkWin() string {
	winPatterns := [8][3][2]int{
		// Rows
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		// Columns
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		// Diagonals
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, pattern := range winPatterns {
		a, b1, c := pattern[0], pattern[1], pattern[2]
		symbol := b.Cells[a[0]][a[1]]
		if symbol != Empty &&
			symbol == b.Cells[b1[0]][b1[1]] &&
			symbol == b.Cells[c[0]][c[1]] {
			return symbol
		}
	}
	return ""
}