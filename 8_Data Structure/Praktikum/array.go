package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make([]string, 0)
	nameMap := make(map[string]bool)

	for _, name := range arrayA {
		if !nameMap[name] {
			nameMap[name] = true
			merged = append(merged, name)
		}
	}

	for _, name := range arrayB {
		if !nameMap[name] {
			nameMap[name] = true
			merged = append(merged, name)
		}
	}

	return merged
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
