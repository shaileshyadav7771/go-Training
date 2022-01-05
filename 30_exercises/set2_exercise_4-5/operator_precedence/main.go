// File name: ...\s02\exercises\set-2\operator_precedence\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: What is the result of the following expressions?
// fmt.Println(4.5*4/2 - 3.5)
// fmt.Printf("%f\n", 4.5*4/2-3.5)

func main() {
	fmt.Print("4.5 * 4 / 2 - 3.5 is: ")
	fmt.Println(4.5*4/2 - 3.5) // 4.5 * 4 / 2 - 3.5 is: 5.5

	// 4.5 * 4 / 2 - 3.5 is: 5.500000
	fmt.Printf("4.5 * 4 / 2 - 3.5 is: %f\n", 4.5*4/2-3.5)
}
