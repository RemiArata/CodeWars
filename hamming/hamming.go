package main

import "fmt"

func main() {
	dna1 := "GAGCCTACTAACGGGAT"
	dna2 := "CATCGTAATGACGGCCT"

	distance := hammingDistance(dna1, dna2)

	fmt.Printf("dna1: %v\n", dna1)
	fmt.Printf("dna2: %v\n", dna2)
	fmt.Printf("The hamming distance between these two strands is: %v", distance)

}

func hammingDistance(dna1, dna2 string) int {
	hamming := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] != dna2[i] {
			hamming++
		}
	}
	return hamming
}
