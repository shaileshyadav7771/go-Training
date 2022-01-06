// File name: ...\s03\exercises\set-5\spinning_cursor1
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

/*
Exercise: Write a program to simulate a spinning cursor.
Press CTRL + C to exit the program.
*/

func main() {
	var delayTime = 20 * time.Millisecond

	fmt.Println("Press CTRL + C to exit the program.")
	fmt.Print("Operation in progress ... ")

	for {
		fmt.Print(`\`)
		time.Sleep(delayTime)
		fmt.Print("\b")

		fmt.Print(`|`)
		time.Sleep(delayTime)
		fmt.Print("\b")

		fmt.Print(`/`)
		time.Sleep(delayTime)
		fmt.Print("\b")

		fmt.Print(`-`)
		time.Sleep(delayTime)
		fmt.Print("\b")
	}
}
