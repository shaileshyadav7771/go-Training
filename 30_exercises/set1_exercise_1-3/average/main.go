// File name: ...\s02\exercises\set-1\average\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Calculate average of three integers

func main() {

	num1 := 12
	num2 := 14
	num3 := 10

	fmt.Printf("average of %d, %d, and %d is %d.\n",
		num1, num2, num3, (num1+num2+num3)/3)
}
