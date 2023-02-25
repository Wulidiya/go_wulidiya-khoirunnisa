package main

import "fmt"

func power(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := power(x, n/2)
	if n%2 == 0 {
		return res * res
	}
	return res * res * x
}

func main() {
	fmt.Println(power(2, 3))
	fmt.Println(power(5, 3))
	fmt.Println(power(10, 2))
	fmt.Println(power(2, 5))
	fmt.Println(power(7, 3))
}
