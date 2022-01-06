// File name: ...\s02\exercises\set-5\utf8-2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Using the for range loop, append the following
// two strings and save them is a slice of 'rune'.
// "𠜎ち" and "再见"
// Also print the result in form of a slice rune and also as a string.
//

func main() {

	s1 := "𠜎ち"
	s2 := "再见"
	var myRunes []rune

	for _, r := range s1 {
		myRunes = append(myRunes, r)
	}
	for _, r := range s2 {
		myRunes = append(myRunes, r)
	}

	// ['𠜎' 'ち' '再' '见']
	fmt.Printf("%q\n", myRunes)

	// "𠜎ち再见"
	fmt.Printf("%q\n", string(myRunes))
}
