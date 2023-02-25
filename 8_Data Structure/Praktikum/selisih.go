package main

import "fmt"

func main() {
	jumlahDiagonalx := 0
	jumlahDiagonaly := 0
	matriks := [3][3]int{
		{1, 2, 3}, // 0
		{4, 5, 6}, // 1
		{9, 8, 9}, // 2
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				jumlahDiagonalx += matriks[i][j]
			}
			if i+j == 2 {
				jumlahDiagonaly += matriks[i][j]
			}
		}
	}
	diff := jumlahDiagonalx - jumlahDiagonaly
	if diff < 0 {
		diff = -diff
	}
	fmt.Println(diff)
}
