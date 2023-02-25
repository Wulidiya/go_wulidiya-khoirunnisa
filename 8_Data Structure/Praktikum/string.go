package main

import (
	"fmt"
	"strconv"
)

func angkamunculSekali(angka string) []int {
	count := make(map[int]int)

	// menghitung jumlah kemunculan setiap angka
	for _, char := range angka {
		num, _ := strconv.Atoi(string(char))
		count[num]++
	}

	// mengumpulkan angka yang hanya muncul 1 kali
	result := []int{}
	for num, freq := range count {
		if freq == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	fmt.Println(angkamunculSekali("1234123"))
	fmt.Println(angkamunculSekali("76523752"))
	fmt.Println(angkamunculSekali("12345"))
	fmt.Println(angkamunculSekali("1122334455"))
	fmt.Println(angkamunculSekali("0872504"))
}
