// File name: ...\s02\exercises\set-4\iota-2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Using 'iota', generate the following output:
// No=0 False=0 Off=0
// Yes=1 True=1 On=1

func main() {
	const (
		No, False, Off = iota, iota, iota // all 0
		Yes, True, On                     // all 1
	)

	// Output: No=0 False=0 Off=0
	fmt.Printf("No=%d False=%d Off=%d\n",
		No, False, Off)

	// Output: Yes=1 True=1 On=1
	fmt.Printf("Yes=%d True=%d On=%d\n",
		Yes, True, On)
}
