package main

import "fmt"

func main() {
	var nth int
	fmt.Print("What prime would you like to find: ")
	fmt.Scan(&nth)

	p := findPrimes(nth)
	fmt.Printf("%v is the prime you are looking for!", p)
}

func findPrimes(n int) int {
	if n <= 3 {
		return n
	}
	var isPrime bool
	seenPrimes := 3
	prime := 3
	i := 4

	for seenPrimes < n {
		isPrime = true
		for j := 2; j < (i/2)+1; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime {
			prime = i
			seenPrimes++
		}
		i++
	}
	return prime
}
