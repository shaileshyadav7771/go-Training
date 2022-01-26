// File name: ...\s12\12_channel_buffered_capacity\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3)
	fmt.Printf("At Start: capacity:%d length:%d \n", cap(c), len(c))

	c <- "Message 1"
	c <- "Message 2"

	fmt.Printf("After sending message1,2 capacity:%d length:%d \n", cap(c), len(c))

	c <- "Message 3"
	fmt.Printf("After sending message3 capacity:%d length:%d \n", cap(c), len(c))

	time.Sleep(5 * time.Second)
	fmt.Println("Reading one value from channel", <-c)

	time.Sleep(2 * time.Second)
	fmt.Println("Reading 2nd value from channel", <-c)

	fmt.Println("Reading last(3rd) value from channel", <-c)

	//Note : Now after this channel length is zero so we can not read further.
	// fmt.Println(<-c) //ok
	// fmt.Println(<-c) //fatal error: all goroutines are asleep - deadlock!

	fmt.Printf("capacity:%d length:%d \n", cap(c), len(c))
}
