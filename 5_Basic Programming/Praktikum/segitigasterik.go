package main

import "fmt"

func main() {
	var tinggi int

	fmt.Print("Masukkan tinggi dari segitiga: ")
	fmt.Scanln(&tinggi)

	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
