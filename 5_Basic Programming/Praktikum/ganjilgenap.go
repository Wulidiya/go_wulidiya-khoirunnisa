package main

import "fmt"

func main() {

	fmt.Print("Masukkan Bilangan Positif : ")
	var num int
	fmt.Scanf("%d", &num)

	if (num % 2) == 0 {
		fmt.Printf("%d adalah genap", num)
	} else {
		fmt.Printf("%d is ganjil", num)
	}
}
