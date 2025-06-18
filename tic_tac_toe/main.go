package main

import "fmt"

func main() {
	board := newBoard()
	currentPlayer := "X"

	for {
		fmt.Println("\nCurrent board:")
		for _, row := range board.Cells {
			fmt.Println(row)
		}

		var row, col int
		for {
			fmt.Printf("%s's turn. Enter row and column (0-2 each): ", currentPlayer)
			n, err := fmt.Scan(&row, &col)
			if err != nil || n != 2 || row < 0 || row > 2 || col < 0 || col > 2 {
				fmt.Println("Invalid input. Please enter two numbers between 0 and 2, like: 1 2")
				fmt.Scanln() // clear invalid input
				continue
			}
			if board.Cells[row][col] != "" {
				fmt.Println("Cell already taken. Choose another.")
				continue
			}
			break
		}

		board.Cells[row][col] = currentPlayer

		if winner := board.checkWin(); winner != "" {
			fmt.Println("\nFinal board:")
			for _, row := range board.Cells {
				fmt.Println(row)
			}
			fmt.Printf("Winner: %s\n", winner)
			break
		}

		if board.isFull() {
			fmt.Println("\nFinal board:")
			for _, row := range board.Cells {
				fmt.Println(row)
			}
			fmt.Println("It's a tie!")
			break
		}

		// Switch turns
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
