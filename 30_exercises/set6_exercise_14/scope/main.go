// File name: ...\s02\exercises\set-6\scope\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: What is the output of the following program?

var a = 1

func main() {
	{
		fmt.Println("Loc 1 - ", a)
		a := 4
		fmt.Println("Loc 2 - ", a)
	}

	a := 2
	fmt.Println("Loc 3 - ", a)

	a += 10

	{
		a += 50
		fmt.Println("Loc 4 - ", a)
	}
	fmt.Println("Loc 5 - ", a)
	changeA()
}

func changeA() {
	fmt.Println("Loc 6 - ", a)

	{
		a := 100
		fmt.Println("Loc 7 - ", a)
	}
	fmt.Println("Loc 8 - ", a)
}

/* Output
Loc 1 -  1
Loc 2 -  4
Loc 3 -  2
Loc 4 -  62
Loc 5 -  62
Loc 6 -  1
Loc 7 -  100
Loc 8 -  1
*/
