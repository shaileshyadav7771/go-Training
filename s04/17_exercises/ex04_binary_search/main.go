// File name: ...\s04\exercises\set-3\binary_search
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

/*
Exercise: Write a program to search for an integer value
using the "Binary Search".

A requirement for binary search:
- Your list needs to be sorted.

*/

func main() {
	// Your list needs to be sorted
	sortedList := []int{
		11, 12, 13, 14, 15, 16, 17, 18,
	}

	var searchKey int
	fmt.Print("Enter an integer: ")
	fmt.Scanf("%d ", &searchKey)

	start := 0
	end := len(sortedList) - 1

	keyFound := false
	var mid int
	for start <= end {
		mid = start + (end-start)/2
		midValue := sortedList[mid]

		// fmt.Println("* ", start, "* ", mid, " * ", midValue, "* ", end)

		if midValue == searchKey {
			keyFound = true
			break
		} else if midValue > searchKey {
			end = mid - 1 //left half of the list
		} else {
			start = mid + 1 //right  half of the list
		}
	}

	if keyFound {
		fmt.Printf("%d was found at index %d.\n", searchKey, mid)

	} else {
		fmt.Printf("%d was not found.\n", searchKey)
	}
}
