// File name: ...\s02\exercises\set-2\display_table\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Write a program that displays the following table:
// i        i*2     i*3
// ==       ===     ===
// 2        4       6
// 3        6       9
// 4        8       12

func main() {

	i := 2

	fmt.Println("i \t i*2 \t i*3")
	fmt.Println("== \t === \t ===")

	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
	i++
	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
	i++
	fmt.Printf("%d \t %d \t %d\n", i, i*2, i*3)
}
