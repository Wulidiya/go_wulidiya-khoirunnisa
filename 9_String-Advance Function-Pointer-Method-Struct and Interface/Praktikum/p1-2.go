package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  string
	score int
}

func (s *Student) input() {
	fmt.Print("Input student's name: ")
	fmt.Scan(&s.name)
	fmt.Print("Input student's score: ")
	fmt.Scan(&s.score)
}

func main() {
	students := make([]Student, 5)
	for i := 0; i < 5; i++ {
		s := &students[i]
		s.input()
	}

	totalScore := 0
	minScore := math.MaxInt32
	maxScore := math.MinInt32
	minName := ""
	maxName := ""

	for _, s := range students {
		totalScore += s.score
		if s.score < minScore {
			minScore = s.score
			minName = s.name
		}
		if s.score > maxScore {
			maxScore = s.score
			maxName = s.name
		}
	}

	averageScore := float64(totalScore) / 5
	fmt.Printf("average score: %.2f\n", averageScore)
	fmt.Printf("min score of students: %s (%d)\n", minName, minScore)
	fmt.Printf("max score of students: %s (%d)\n", maxName, maxScore)
}
