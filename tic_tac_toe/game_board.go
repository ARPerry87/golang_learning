package main

import "fmt"

type Board struct {
	Cells [3][3]string
}

func newBoard() *Board {
	return &Board{
		Cells: [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}
}

// Display the board
func (b *Board) Print() {
	fmt.Println("\nCurrent board:")
	for _, row := range b.Cells {
		fmt.Println(row)
	}
}

// Ask the player for a valid move
func (b *Board) PromptMove(player string) (int, int) {
	var row, col int
	for {
		fmt.Printf("%s's turn. Enter row and column (0-2 each): ", player)
		n, err := fmt.Scan(&row, &col)
		if err != nil || n != 2 || row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid input. Please enter two numbers between 0 and 2, like: 1 2")
			fmt.Scanln() // flush invalid input
			continue
		}
		if b.Cells[row][col] != "" {
			fmt.Println("Cell already taken. Choose another.")
			continue
		}
		break
	}
	return row, col
}

// Apply a move to the board
func (b *Board) PlaceMove(row, col int, symbol string) {
	b.Cells[row][col] = symbol
}

// Check for three matching non-empty cells
func (b *Board) allEqual(a, b1, c [2]int) string {
	symbol := b.Cells[a[0]][a[1]]
	if symbol != "" &&
		symbol == b.Cells[b1[0]][b1[1]] &&
		symbol == b.Cells[c[0]][c[1]] {
		return symbol
	}
	return ""
}

// Check if there's a winner
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

// Check if the board is full
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
