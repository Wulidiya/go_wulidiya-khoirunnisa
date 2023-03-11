package main

import "fmt"

func Frog(jumps []int) int {
	// Inisialisasi array cost dengan nilai 0
	cost := make([]int, len(jumps))
	cost[0] = 0

	// Perulangan untuk mengisi nilai cost
	for i := 1; i < len(jumps); i++ {
		cost[i] = cost[i-1] + abs(jumps[i]-jumps[i-1])

		if i > 1 {
			cost[i] = min(cost[i], cost[i-2]+abs(jumps[i]-jumps[i-2]))
		}
	}

	// Nilai minimum berada pada elemen terakhir array cost
	return cost[len(cost)-1]
}

// Fungsi untuk menghitung nilai absolut dari suatu bilangan
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Fungsi untuk mencari nilai minimum dari dua bilangan
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
