// File name: ...\s02\exercises\set-5\utf8-1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"unicode/utf8"
)

// Exercise: Which one of the following strings take more space (in bytes)?
// 	"Hello" and "𠜎ち"

func main() {
	s1 := "Hello"
	s2 := "𠜎ち"

	// 'RuneCountInString' returns number of runes (characters in this case)
	// 'len' returns the number of bytes occupied by the string.
	fmt.Println(len(s1), utf8.RuneCountInString(s1)) //5 5
	fmt.Println(len(s2), utf8.RuneCountInString(s2)) //7 2
}
