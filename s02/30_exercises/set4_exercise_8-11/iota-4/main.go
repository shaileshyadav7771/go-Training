// File name: ...\s02\exercises\set-4\iota-4\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: What's result of the following code?
// const (
// 	_      = iota
// 	Ten    = 10 * iota
// 	Twenty = iota
// 	Thirty
// )

func main() {

	const (
		// ignoring the first value by assigning to blank identifier
		_      = iota
		Ten    = 10 * iota
		Twenty = iota
		Thirty
	)

	// Output: Ten=10 Twenty=2 Thirty=3
	fmt.Printf("Ten=%d Twenty=%d Thirty=%d\n",
		Ten, Twenty, Thirty)
}
