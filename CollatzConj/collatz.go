package main

import (
	"fmt"
	"os"
)

func collatz(start int) []int {
	steps := []int{}
	n := start

	for n != 1 {
		steps = append(steps, n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
	steps = append(steps, n)
	return steps
}

func main() {
	var userInput int
	fmt.Print("plz enter a positive number greater than 1: ")
	fmt.Scan(&userInput)
	if userInput < 1 {
		fmt.Printf("%v is not a valid start value. Must be greater than or equal to 1.", userInput)
		os.Exit(0)
	}

	s := collatz(userInput)
	fmt.Printf("The input was %d\n", userInput)
	fmt.Printf("The collatz steps are: %v\n", s)
}
