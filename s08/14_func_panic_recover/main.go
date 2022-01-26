// File name: ...\s08\14_func_panic_recover\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	//we are using recover() function so basically It will capture the Error message and display that error only
	defer func() {
		errMsg := recover()
		fmt.Println("Our custome message -Testing")
		fmt.Println(errMsg)
	}()

	// Only use one of the following at the same time.

	// var x map[string]int
	// //It is empty/nill map and we can't assign value to nill map
	// x["key"] = 10
	// fmt.Println(x)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	panic("BOO-Some Error is coming..!!!")
	//paniic generate RUN time ERROR so code after this line will be unreachable.
	//if we'll run then this recover() function will show panic ""This line won't be reached!""
	fmt.Println("This line won't be reached!")
}
