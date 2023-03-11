package main

import "fmt"

// memoization digunakan untuk menyimpan hasil perhitungan sebelumnya
var memo map[int]int

func fibo(n int) int {
	if n <= 1 {
		return n
	}

	// jika sudah pernah dihitung sebelumnya, kembalikan hasilnya dari memoization
	if val, ok := memo[n]; ok {
		return val
	}

	// Fibonacci(n) dihitung menggunakan rekursi top-down
	result := fibo(n-1) + fibo(n-2)

	// hasil perhitungan disimpan ke dalam memoization
	memo[n] = result

	return result
}

func main() {
	memo = make(map[int]int)
	fmt.Println(fibo(0))
	fmt.Println(fibo(1))
	fmt.Println(fibo(2))
	fmt.Println(fibo(3))
	fmt.Println(fibo(5))
	fmt.Println(fibo(6))
	fmt.Println(fibo(7))
	fmt.Println(fibo(9))
	fmt.Println(fibo(10))
}
