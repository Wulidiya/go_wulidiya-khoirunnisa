package main

import (
	"fmt"
)

func printMultiples(ch chan<- int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
}

func main() {
	ch := make(chan int, 10) // Buffer channel dengan kapasitas 10

	go printMultiples(ch)

	// Menerima data dari channel dan mencetaknya
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
