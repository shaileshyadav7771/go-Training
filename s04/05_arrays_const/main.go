// File name: ...\s04\05_arrays_const\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	type Size int

	const (
		EX Size = iota
		LG
		MD
		SM
	)

	//same thing we can define in another way like below so basically It is same thing's
	//Above we are using just to see how we can use type with an Array :)
	// const (
	// 	EX = iota
	// 	LG
	// 	MD
	// 	SM
	// )

	sz := [...]string{EX: "Extra Large", LG: "Large", MD: "Medium", SM: "Small"}
	fmt.Println(MD, sz[MD])
	fmt.Println(sz)

	x := [...]int{3: 10, 20}
	fmt.Println(x)
}
