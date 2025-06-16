package main

import "fmt"

func isEvenOrOdd() []int{
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range slice {
		if v%2 == 0 {
			fmt.Printf("%d is even\n", v)
		} else {
			fmt.Printf("%d is odd\n", v)
		}
	}
	return slice
}
