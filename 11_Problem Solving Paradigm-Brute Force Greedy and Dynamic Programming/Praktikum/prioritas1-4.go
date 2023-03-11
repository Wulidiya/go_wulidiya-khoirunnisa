package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	found := false
	for x := 1; x <= b; x++ {
		if b%x != 0 {
			continue
		}
		for y := x + 1; y <= b/x; y++ {
			if b%(x*y) != 0 {
				continue
			}
			z := b / (x * y)
			if x+y+z == a && x*x+y*y+z*z == c {
				fmt.Printf("%d %d %d\n", x, y, z)
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3)
	SimpleEquations(6, 6, 14)
}
