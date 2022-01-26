// File name: ...\s08\12_func_defer_sqaure\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT - Use the 'defer' statement to write a square function that hijacks the
// correct return type when the input parameter is 2 or 4.
// For 2 => return 6; For 4 => return 20;   formula: (x*x)+x
// For other values return (x*x)

func main() {
	fmt.Println(square(2))
	fmt.Println(square(4))
	fmt.Println(square(10))

}

func square(x int) (result int) {

	result = x * x

	defer func() {
		if x == 2 || x == 4 {
			result += x
		}

	}()

	fmt.Println("Just test purpose here This line will excecute before defer statement.")
	return

}
