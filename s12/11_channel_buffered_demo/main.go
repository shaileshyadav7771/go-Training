// File name: ...\s12\11_channel_buffered_demo\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // test with 1 to 3

	for i := 1; i <= 3; i++ {
		go printMsg(c, i)
	}
	// time.Sleep(10000 * time.Second)
	var input string
	fmt.Scanln(&input) // press "Enter" to exit
}

func printMsg(c chan int, id int) {
	fmt.Printf("Hey %d is waiting for a channel space...\n", id)

	c <- id
	fmt.Printf("+++ %d has acquired channel space\n", id)

	time.Sleep(1 * time.Millisecond)
	fmt.Printf("--- %d has released the channel space\n", id)

	<-c
}
