package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float64

	fmt.Print("Masukkan sisi Atas: ")
	fmt.Scanln(&sisiAtas)

	fmt.Print("Masukkan sisi Bawah: ")
	fmt.Scanln(&sisiBawah)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	luas := (sisiAtas + sisiBawah) * tinggi / 2

	fmt.Printf("Jadi, luas trapesium adalah: %.2f\n", luas)
}
