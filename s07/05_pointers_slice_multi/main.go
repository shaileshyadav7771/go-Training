// File name: ...\s07\05_pointers_slice_multi\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	slices := [][]int{{2, 3, 4}, {5, 6}, {7, 8, 9, 10}}

	fmt.Println("Slice of slice :", slices) //Slice of slice : [[2 3 4] [5 6] [7 8 9 10]]

	var bigSlice []int

	var p *[][]int

	p = &slices

	for r := range *p {
		bigSlice = append(bigSlice, (*p)[r]...)
		//... lapses used to convert a slice into a list of element inorder to add it
		// fmt.Printf("* %d %v %d\n", r, (*p)[r], (*p)[r][0])

	}

	fmt.Printf("%v %T\n", *p, *p)
	fmt.Printf("%v %T\n", slices, slices)
	fmt.Printf("%v %T\n", bigSlice, bigSlice)
}
