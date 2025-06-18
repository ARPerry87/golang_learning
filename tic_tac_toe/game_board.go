package main

type Board struct {
	Cells [3][3]string
}

// Create a new empty board
func newBoard() *Board {
	return &Board{
		Cells: [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}
}

// Check if three coordinates all have the same non-empty symbol
func (b *Board) allEqual(a, b1, c [2]int) string {
	symbol := b.Cells[a[0]][a[1]]
	if symbol != "" &&
		symbol == b.Cells[b1[0]][b1[1]] &&
		symbol == b.Cells[c[0]][c[1]] {
		return symbol
	}
	return ""
}

// Check for a winner on the board
func (b *Board) checkWin() string {
	// Rows
	for i := range 3 {
		if winner := b.allEqual([2]int{i, 0}, [2]int{i, 1}, [2]int{i, 2}); winner != "" {
			return winner
		}
	}
	// Columns
	for j := range 3 {
		if winner := b.allEqual([2]int{0, j}, [2]int{1, j}, [2]int{2, j}); winner != "" {
			return winner
		}
	}
	// Diagonals
	if winner := b.allEqual([2]int{0, 0}, [2]int{1, 1}, [2]int{2, 2}); winner != "" {
		return winner
	}
	if winner := b.allEqual([2]int{0, 2}, [2]int{1, 1}, [2]int{2, 0}); winner != "" {
		return winner
	}
	return ""
}

// Check if the board is full (tie condition)
func (b *Board) isFull() bool {
	for i := range 3 {
		for j := range 3 {
			if b.Cells[i][j] == "" {
				return false
			}
		}
	}
	return true
}
