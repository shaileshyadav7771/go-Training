// File name: ...\s04\16_slices_operations\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: implementing basic operations with slice.
// Use the slicing techniques you've learned so far to implement
// the following basic operations with slices:
// Append (a=[1 2 3] & b=[4 5 6 7] -> c=[1 2 3 4 5 6 7])
// Copy (a=[1 2 3] -> b=[1 2 3])
// Cut (a=[1 2 3 4 5 6 7] -> a=[1 2 5 6 7])
// Delete (with preserving order)    (a=[1 2 5 6 7] -> a=[1 5 6 7])
// Delete (without preserving order) (a=[1 2 3 4 5 6 7] -> a=[1 2 7 4 5 6])
// Expand (a=[1 2 7 4 5 6] -> a=[1 2 0 0 0 0 7 4 5 6])
// Extend (a=[1 2 0 0 0 0 7 4 5 6] -> a=[1 2 0 0 0 0 7 4 5 6 0 0 0])
// Insert (a=[1 2 0 0 0 0 7 4 5 6 0 0 0] -> a=[1 2 0 9 10 0 0 0 7 4 5 6 0 0 0])
//
// Ref: https://github.com/golang/go/wiki/SliceTricks
//
func main() {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}

	a = append(a, b...) // Append
	fmt.Printf(" (1) a%v b%v \n", a, b)

	c := make([]int, len(a))
	copy(c, a) // Copy
	fmt.Printf(" (2) a%v c%v \n", a, c)

	d := append([]int(nil), a...) // Copy
	fmt.Printf(" (3) a%v d%v \n", a, d)

	c = append(c[:2], c[4:]...) // Cut
	fmt.Printf(" (4) c%v \n", c)

	c = append(c[:1], c[2:]...) // Delete (with preserving order)
	fmt.Printf(" (5) c%v \n", c)

	c = c[:1+copy(c[1:], c[2:])] // Delete (with preserving order)
	fmt.Printf(" (6) c%v \n", c)

	// Delete (without preserving order)
	fmt.Printf(" (7) d%v \n", d)
	d[2] = d[len(d)-1]
	d = d[:len(d)-1]
	fmt.Printf(" (8) d%v \n", d)

	d = append(d[:2], append(make([]int, 4), d[2:]...)...) // Expand
	fmt.Printf(" (9) d%v \n", d)

	d = append(d, make([]int, 3)...) // Extend
	fmt.Printf("(10) d%v \n", d)

	d = append(d[:3], append([]int{9, 10}, d[3:]...)...) // Insert
	fmt.Printf("(11) d%v \n", d)

	testCase1()
	testCase2()
}

func testCase1() {
	fmt.Println()
	x := []int{1, 5, 6, 7}

	x1 := x[1:] //[5 6 7]
	x2 := x[2:] //[6 7]
	x3 := x[:1] //[1]

	fmt.Printf("\nx=%v, x1=%v x2=%v x3=%v", x, x1, x2, x3)

	// don't forget that slices are not actual copies.
	// They point to the base array, therefore the following
	// copy is changing the base array by copying 6,7 at
	// index 2 (as x2 suggests) to the index that x1
	// suggests (that is index 1) and that's why 5,6
	// change to 6,7 (index 1 and 2).
	x4 := copy(x1, x2)
	fmt.Printf("\nx=%v, x1=%v x2=%v x3=%v x4=%v", x, x1, x2, x3, x4)

	// from the previous line x4 returned 2, therefore the
	// following is like x5 := x[:3] , meaning from index 0 to 2 (inclusive)
	x5 := x[:1+x4]
	fmt.Printf("\n\nx5=%v", x5)

	fmt.Println()
}
func testCase2() {
	fmt.Println("\n")

	d := []int{1, 2, 3, 4, 5, 6, 7}
	i := len(d) - 1 //number of of elements od d, minus 1
	// fmt.Printf("i=%v i=%v\n", i, d[i]) //6

	fmt.Printf("i=%v, d=%v d[2]=%v d[i]=%v\n", i, d, d[2], d[i])

	// d[i] is like d[len(d)-1]
	d[2] = d[i]
	fmt.Printf("i=%v, d=%v d[2]=%v d[i]=%v\n", i, d, d[2], d[i])
}
