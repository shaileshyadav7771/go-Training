// File name: ...\s02\exercises\set-4\iota-1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: What's the result of the following code:
// const (
// 	  c1 = iota
// 	  c2 = 5
// 	  c3 = iota
// 	  c4 = 8
// 	  c5 = iota
// )

func main() {
	const (
		c1 = iota // 0
		c2 = 5    // 5
		c3 = iota // 2
		c4 = 8    // 8
		c5 = iota // 4
	)

	// output: c1=0 c2=5 c3=2 c4=8 c5=4
	fmt.Printf("c1=%d c2=%d c3=%d c4=%d c5=%d\n",
		c1, c2, c3, c4, c5)

}
