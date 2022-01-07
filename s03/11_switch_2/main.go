// File name: ...\s03\11_switch_2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

//fallthrough -when mentioned it'll not check the expression for next swich case and enter into block.
package main

import (
	"fmt"
)

func main() {
	var seasonNo int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &seasonNo)

	switch seasonNo {
	case 1:
		fmt.Println("spring - ", seasonNo)
	case 2:
		fmt.Println("summer - ", seasonNo)
		fallthrough
	case 3:
		fmt.Println("fall - ", seasonNo)
		fallthrough
	case 4:
		fmt.Println("winter - ", seasonNo)
		fallthrough
	default:
		fmt.Println("a new season - ", seasonNo)
		// fallthrough                             // fallthrough we can't define in default statment :)
	}
}
