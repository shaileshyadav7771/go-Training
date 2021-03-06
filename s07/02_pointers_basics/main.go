// File name: ...\s07\02_pointers_basics\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	x := 10

	p := &x
	// var p *int = &x 				// 2nd way way-redundant
	// var p *int; p = &x			// 3rd way-redundant

	fmt.Printf("x=%T &x=%T p=%T *p=%T &p=%T \n\n", x, &x, p, *p, &p)

	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n", x, &x, p, *p, &p)

	*p = 2 // dereferencing
	fmt.Printf("Dereferencing to 2-: x=%d &x=%x p=%x *p=%d &p=%x \n", x, &x, p, *p, &p)

	x = 3
	fmt.Printf("x=%d &x=%x p=%x *p=%d &p=%x \n", x, &x, p, *p, &p)

	y := 4
	p = &y
	fmt.Printf("y=%d &y=%x p=%x *p=%d &p=%x \n", y, &y, p, *p, &p)

	q := new(int)

	fmt.Println()
	fmt.Printf("*q=%d q=%x \n", *q, q)

	*q = 10
	fmt.Printf("*q=%d q=%x \n", *q, q)

	//Note:%x for formatting 16 digit no.

	// Note: if we run multiple time we may see the Address is changing(It is because of
	//Memory Allocation.)

	//Note: &x and p both are pointing to the same address. (check notes for more details.)
}
