package main

import "fmt"

func MaxSequence(arr []int) int {
	maxSum, currentSum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		currentSum = Max(currentSum+arr[i], arr[i])
		maxSum = Max(currentSum, maxSum)
	}
	return maxSum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
