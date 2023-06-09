package main

import (
	"fmt"
	"sync"
)

var (
	factorial      = 1
	factorialMutex sync.Mutex
	wg             sync.WaitGroup
)

func calculateFactorial(n int) {
	defer wg.Done()

	factorialMutex.Lock()
	defer factorialMutex.Unlock()

	for i := 1; i <= n; i++ {
		factorial *= i
	}
}

func main() {
	n := 5 // Angka untuk perhitungan faktorial

	wg.Add(1)
	go calculateFactorial(n)

	wg.Wait()

	fmt.Printf("Faktorial dari %d adalah %d\n", n, factorial)
}
