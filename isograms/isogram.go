package main

import "fmt"

func main() {
	var s string

	fmt.Print("enter a word: ")
	fmt.Scan(&s)
	b := isogram(s)
	if b {
		fmt.Printf("%v is a isogram!", s)
	} else {
		fmt.Printf("%v is not an isogram :(", s)
	}
}

func isogram(s string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			return false
		} else {
			m[s[i]] = 1
		}
	}
	return true
}
