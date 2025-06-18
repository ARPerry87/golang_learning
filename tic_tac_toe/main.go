package main

import "fmt"

func main() {
	board := newBoard()
	currentPlayer := "X"

	for {
		board.Print()

		row, col := board.PromptMove(currentPlayer)
		board.PlaceMove(row, col, currentPlayer)

		if winner := board.checkWin(); winner != "" {
			board.Print()
			fmt.Printf("Winner: %s\n", winner)
			break
		}

		if board.isFull() {
			board.Print()
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
