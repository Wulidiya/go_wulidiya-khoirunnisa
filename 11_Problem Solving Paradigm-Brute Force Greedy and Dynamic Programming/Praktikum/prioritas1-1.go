package main

import "fmt"

func bilanganbiner(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("%b ", i)
	}
}

func main() {
	bilanganbiner(3)
}
