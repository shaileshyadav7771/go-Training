// File name: ...\s08\01_func_literal\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	printMsg("Calling a function!")

	// It is nameless Function
	showMsg := func(msg string) {
		fmt.Println(msg)
	}

	showMsg("Using a function literal!")
	fmt.Printf("%T\n", showMsg)

	func(msg string) {
		fmt.Println(msg)
	}("quickly reacting!")

	func(msg string) {
		fmt.Println(msg)
	}("Shailesh")

	func(msg int) {
		fmt.Println(msg)
	}(5 + 5)
}

func printMsg(msg string) {
	fmt.Println(msg)
}

//Note: As we know if we are declaring/Creating a seperate function then It will take some memory storage and
//It'll be available till called within that function out of fn that variable will be pop out. So If we declaring
//nameless fn obesivoly It'll help us in low level of programming like networking application's.
