// File name: ...\s07\06_pointers_func\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	//Testing Purpose
	// var a int = 10
	// fmt.Println("&x ki value is: %x", &a)
	// fmt.Println()

	var p1 = f(2)
	fmt.Printf("p1=%x *p1=%d &p1=%x\n\n", p1, *p1, &p1)

	//Note: Here we are able to access *p1 which is coming from another stack f (but before this we have
	//discussed that local variable decalred inside a functon/stack is limited to that only menas once
	//it is over then that value will not be accessible but here It is happening opposite in case of pointer.)

	var p2 = f(3)
	fmt.Printf("p2=%x *p2=%d &p2=%x\n", p2, *p2, &p2)
}

//This is the below fn which'll return pointer to the pass i/p int value
func f(inp int) *int {
	v := inp * 2
	fmt.Printf("&v=%x\n", &v)

	return &v
}
