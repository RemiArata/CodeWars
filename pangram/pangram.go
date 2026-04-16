package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	testData := "The quick brown fox jumps over the lazy dog"

	b := pangram(testData)
	if b {
		fmt.Println("The string is a panogram")
	} else {
		fmt.Println("The string is not a panogram ")
	}
}

func pangram(s string) bool {
	sLower := strings.ReplaceAll(strings.ToLower(s), " ", "")
	m := make(map[byte]int)

	for i := 0; i < len(sLower); i++ {
		if _, ok := m[sLower[i]]; ok {
			m[sLower[i]]++
		} else {
			m[sLower[i]] = 0
		}
	}
	keySlice := slices.Collect(maps.Keys(m))
	return len(keySlice) == 26
}
