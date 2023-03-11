package main

import "fmt"

func toRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}
	return roman
}

func main() {
	fmt.Println(toRoman(4))
	fmt.Println(toRoman(9))
	fmt.Println(toRoman(23))
	fmt.Println(toRoman(2021))
	fmt.Println(toRoman(1646))
}
