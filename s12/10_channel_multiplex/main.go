// File name: ...\s12\10_channel_multiplex\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	stats := make(map[string]int)
	fmt.Println(stats)

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	for i := 0; i < 20; i++ {

		go func() {
			time.Sleep(10 * time.Second)
			c1 <- "Hello from customer service #1"
		}()

		go func() {
			time.Sleep(8 * time.Second)
			c2 <- "Hello from customer service #2"
		}()

		go func() {
			time.Sleep(6 * time.Second)
			// fmt.Println("stats", stats)
			c3 <- "Hello from customer service #3"

		}()

		select {

		case msg1 := <-c1:
			stats["Shailesh"]++
			time.Sleep(time.Second)
			fmt.Println(msg1)
		case msg2 := <-c2:
			stats["Munna"]++
			time.Sleep(time.Second)
			fmt.Println(msg2)
		case msg3 := <-c3:
			stats["Akhilesh"]++
			time.Sleep(time.Second)
			fmt.Println(msg3)
		default:
			stats["Customer Waiting"]++
			time.Sleep(2 * time.Second)
			fmt.Println("No customer service is available at this time!")

		}

	}

	fmt.Printf("\n***Customer Service Stats***\n%v", stats)

	close(c1)
	close(c2)
	close(c3)
}
