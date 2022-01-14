// File name: ...\s04\10_slices_append_ellipsis\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	var f []float32
	fmt.Print("Value of f is:", f, "\n")

	f = append(f, 1.2)
	fmt.Print("Value of f is:", f, "\n")

	f = append(f, 2.4, 4.8, 8.16)

	fmt.Println(f)

	//imp f is slice and we want to add array that is why we are adding array and using f...

	f = append(f, f...)
	fmt.Println(f)
	fmt.Printf("%T,%v", f, f)
}
