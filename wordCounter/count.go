package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var data string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Plz enter a string to count the words in\n=>")
	if scanner.Scan() {
		data = scanner.Text()
	}

	sanitizedData := sanitizeInput(data)
	r := wordCount(sanitizedData)
	fmt.Println(r)
}

func sanitizeInput(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\v", "")
	s = strings.ReplaceAll(s, "\\", "")
	s = strings.ReplaceAll(s, "\"", "")
	re := regexp.MustCompile(`\p{p}`)
	s = re.ReplaceAllString(s, "")
	return s
}

func wordCount(s string) map[string]int {
	splitString := strings.Split(s, " ")
	wrdCounts := make(map[string]int)

	for _, wrd := range splitString {
		if wrd != "" {
			if _, ok := wrdCounts[wrd]; ok {
				wrdCounts[wrd]++
			} else {
				wrdCounts[wrd] = 1
			}
		}
	}
	return wrdCounts
}
