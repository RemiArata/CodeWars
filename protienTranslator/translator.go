package main

import (
	"fmt"
	"strings"
)

func main() {
	testData := "AUGUUUUCUUAAAUG"
	r := translator(testData)

	fmt.Printf("Input data: %v\n", testData)
	fmt.Printf("Resulting:  %v\n", r)

}

func translator(t string) string {
	var strAcc string
	sol := []string{}
	t = strings.ToUpper(t)
	protiens := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAC": "Tyrosine",
		"UAU": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
	}

	for _, chr := range t {
		strAcc += string(chr)
		if len(strAcc) == 3 {
			fmt.Println(strAcc)
			if strAcc == "UAA" || strAcc == "UAG" || strAcc == "UGA" {
				return strings.Join(sol, " ")
			}
			sol = append(sol, protiens[strAcc])
			strAcc = ""
		}

	}
	return strings.Join(sol, " ")
}
