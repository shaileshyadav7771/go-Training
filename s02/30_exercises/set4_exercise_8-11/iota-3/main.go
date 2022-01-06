// File name: ...\s02\exercises\set-4\iota-3\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Generate the following output using 'iota':
// Ten=10 Twenty=20 Thirty=30

func main() {
	type ByteSize float64

	const (
		// ignoring the first value by assigning to blank identifier
		_   = iota
		Ten = 10 * iota
		Twenty
		Thirty
	)

	// Output: Ten=10 Twenty=20 Thirty=30
	fmt.Printf("Ten=%d Twenty=%d Thirty=%d\n",
		Ten, Twenty, Thirty)
}
