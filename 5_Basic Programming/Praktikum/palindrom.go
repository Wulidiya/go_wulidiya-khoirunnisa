package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Masukkan satu kata: ")
	var input string
	fmt.Scanln(&input)

	if isPalindrome(input) {
		fmt.Printf("%s adalah palindrome\n", input)
	} else {
		fmt.Printf("%s adalah bukan palindrome\n", input)
	}
}
func isPalindrome(input string) bool {
	input = strings.ToLower(input)
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
