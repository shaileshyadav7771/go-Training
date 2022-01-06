// File name: ...\s02\08_string\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	fmt.Println(len("Good Day"))
	fmt.Println("Good Day"[5])
	fmt.Println()

	s := "SuperHelpful"

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:9])
	fmt.Println(s[:])

	fmt.Println(s[:5] + s[5:])

	//s[0] = 'D' // compile error:
	//here we are getting an error because string is immutable (same in Python we saw..)

	fmt.Println("s value address :", &s)
	s += "Friend"
	y := s
	fmt.Println(s)
	//string is immutable so we can't modify so for above go will stote at new memory place.
	fmt.Println("s value address after concatenation", &y)

	a := "Shailesh"
	fmt.Println("step size eq=2: ", a[1:7])
}
