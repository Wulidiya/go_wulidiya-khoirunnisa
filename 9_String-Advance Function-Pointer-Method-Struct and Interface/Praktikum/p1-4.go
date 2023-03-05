package main

import "fmt"

func MinMaX(numbers ...*int) (min, max int) {
	min = *numbers[0]
	// max = *numbers[0]
	for _, angka := range numbers {
		if *angka < min {
			min = *angka
		}

		if *angka > max {
			max = *angka
		}

	}
	for i := 0; i < len(numbers); i++ {
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6)

	min, max = MinMaX(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai minimum:", min)
	fmt.Println("Nilai maksimum:", max)

}
