// File name: ...\s08\08_func_recursion_factorial\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Println("=>", factorial(30))
	fmt.Println("=>", factorial(60))
	fmt.Println("=>", factorial(1000))
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	} else {
		// fmt.Print(n, " ")
		return n * factorial(n-1)
	}

}

// func factorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	fmt.Print(n, " ") //for debugging purposes
// 	return n * factorial(n-1)
// }
