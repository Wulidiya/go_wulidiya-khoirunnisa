package main

import (
	"fmt"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	if s.nameEncode != "" {
		return s.nameEncode
	}

	var nameEncode = cipherSubstitute(s.name)
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	if s.nameEncode == "" {
		return s.name
	}

	var nameDecode = cipherSubstitute(s.nameEncode)
	return nameDecode
}

func cipherSubstitute(s string) string {
	result := []rune{}
	for _, val := range s {
		if val == ' ' {
			result = append(result, ' ')
			continue
		}
		result = append(result, 'z'-((val)-'a'))
	}
	return string(result)
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name "+a.name+" is: ", c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Encrypted Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student’s name "+a.nameEncode+" is: ", c.Decode())
	}
}
