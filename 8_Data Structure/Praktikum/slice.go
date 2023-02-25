package main

import "fmt"

func Mapping(slice []string) map[string]int {
	m := make(map[string]int)
	for _, s := range slice {
		m[s]++
	}
	return m
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
