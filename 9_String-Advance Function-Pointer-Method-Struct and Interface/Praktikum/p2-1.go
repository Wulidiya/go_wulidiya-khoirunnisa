package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	for _, char := range input {
		// Mengecek jika karakter adalah huruf kecil
		if char >= 'a' && char <= 'z' {
			// Melakukan pergeseran karakter
			shiftedChar := ((int(char)-'a')+offset)%26 + 'a'
			// Menambahkan karakter hasil pergeseran ke dalam hasil
			result += string(shiftedChar)
		} else {
			// Menambahkan karakter non-huruf ke dalam hasil
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
