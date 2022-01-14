// File name: ...\s08\06_func_callback\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {

	square := func(i int) int {
		return i * i
	}

	cubes := func(i int) int {
		return i * i * i
	}

	fmt.Printf("%v\n", calc(square, 3))
	fmt.Printf("%v\n", calc(cubes, 3))

	// Can you tell what techniques have been used in the following block?
	//Ans: We are using both the things Ananonmyous/Nameless function and callback(above 2 line we have used)
	fmt.Printf("%v\n", calc(func(i int) int {
		return i * i
	}, 3))
}

//so here what is happening is first calc(square, 3) so this square assign to >f of below calc fn and value 3
//will be assign to > x int
func calc(f func(int) int, x int) int {
	return f(x)
}
