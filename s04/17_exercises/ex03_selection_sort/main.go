// File name: ...\s04\exercises\set-4\selection_sort
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

/*
Exercise: Write a program to sort a list of
integer values using the 'Selection Sort'.

Logic:
Select the smallest number and swap it with the first number.
Select the smallest number and swap it with the second number...

Example:
18, 20, 14, 17, 13
13, 20, 14, 17, 18
13, 14, 20, 17, 18
13, 14, 17, 20, 18
13, 14, 17, 18, 20

*/

func main() {
	items := []int{
		18, 20, 14, 17, 13,
		16, 11, 15, 12, 19,
	}

	end := len(items)

	for i := 0; i < end-1; i++ {
		minIdx := i
		for j := i + 1; j < end; j++ {
			if items[minIdx] > items[j] {
				minIdx = j
			}
		}

		if items[minIdx] < items[i] {
			items[minIdx], items[i] = items[i], items[minIdx]
		}
	}

	fmt.Printf("%v\n", items)
}
