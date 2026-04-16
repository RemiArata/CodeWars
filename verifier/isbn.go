package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// testData := "0-306-40615-2"
	// testData := "0-13-149505-X"
	testData := "007462542X"
	b := validISBN(testData)
	fmt.Print(b)
}

func validISBN(num string) bool {
	num = strings.ReplaceAll(num, "-", "")
	if len(num) != 10 {
		return false
	}
	var sum int

	for i, dig := range num {
		if string(dig) == "X" && i == 9 {
			sum += 10
		} else {
			n, ok := strconv.Atoi(string(dig))
			if ok != nil {
				fmt.Println("Invalid digit in ISBN number. Halting execution.")
				os.Exit(1)
			}
			sum += (10 - i) * int(n)
		}
	}
	return sum%11 == 0

}
