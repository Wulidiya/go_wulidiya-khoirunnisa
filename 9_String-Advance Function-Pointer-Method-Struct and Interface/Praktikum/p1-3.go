package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var result string

	// Mencari string yang memiliki panjang lebih pendek
	shorter := a
	longer := b
	if len(a) > len(b) {
		shorter = b
		longer = a
	}

	// Mencari substring dengan panjang terpanjang yang sama
	for i := 0; i < len(shorter); i++ {
		for j := len(shorter) - i; j > 0; j-- {
			substr := shorter[i : i+j]
			if strings.Contains(longer, substr) {
				if len(substr) > len(result) {
					result = substr
				}
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
