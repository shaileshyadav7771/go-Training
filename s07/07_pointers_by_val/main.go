// File name: ...\s07\07_pointers_by_val\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	x := 10

	fmt.Printf("x=%d\n", x)
	f1(x)
	fmt.Printf("x=%d\n", x)
}

//Note:: Here we are sending a copy of s to f1 and not the reference of it :).
//In other programming lang It's called call by value
func f1(y int) {
	fmt.Printf("(f1) y=%d\n", y)
	y += 5
	fmt.Printf("(f1) y=%d\n", y)
}
