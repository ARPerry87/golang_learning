package main

import "fmt"

func main() {
    board := filledBoard()
    fmt.Println("Filled Board:")
    for _, row := range board.Cells {
        fmt.Println(row)
    }

    winner := board.checkWin()
    if winner != "" {
        fmt.Printf("Winner: %s\n", winner)
    } else {
        fmt.Println("No winner.")
    }
}
