// File name: ...\s03\exercises\set-8\mult_table
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

/*
Exercise: Write a program to create a multiplication table.
Design your code so that it can easily adjust to display
tables with different lengths. Make sure that column headers
will dynamically adjust accordingly.

Example of a 2x3 table
        1  2  3
- - - - -  -  -
  1 |   1  2  3
  2 |   2  4  6

Example of a 3x2 table
        1  2
- - - - -  -
  1 |   1  2
  2 |   2  4
  3 |   3  6

*/

func main() {
	const numOfRows = 3
	const numOfCols = 6

	fmt.Printf("\n%6s", " ")
	for j := 1; j <= numOfCols; j++ {
		fmt.Printf("%3d", j)
	}

	fmt.Print("\n- - - -")
	for j := 1; j <= numOfCols; j++ {
		fmt.Printf("%2s ", "-")
	}

	for i := 1; i <= numOfRows; i++ {
		fmt.Printf("\n%3d | ", i)
		for j := 1; j <= numOfCols; j++ {
			fmt.Printf("%3d", i*j)
		}
	}

	fmt.Println()
}
