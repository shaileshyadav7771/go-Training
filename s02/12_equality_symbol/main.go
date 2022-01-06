// File name: ...\s02\12_equality_symbol\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	s := "hi"
	fmt.Println("s variable memory add:", &s)
	s = "hello"
	fmt.Println("s variable memory add:", &s)

	fmt.Println(s == "hello")
	fmt.Println(s != "hello")

	var (
		name1 = "Tim"
		name2 = "Ava"
		name3 = "Nick"
	)

	fmt.Println(name1, name2, name3)

	var y string = "testing"
	fmt.Println("y variable memory add:", &y, y)
	z := y
	fmt.Println("z variable memory add:", &z, z)

}
