// File name: ...\s12\04_channel_unbuffered\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sync"
)

var waitG sync.WaitGroup

func main() {
	c := make(chan int)

	waitG.Add(1)
	go send(c)
	go receive(c)
	waitG.Wait()

	// time.Sleep(2 * time.Second)
	// time.Sleep(3 * time.Millisecond)

}

func send(c chan int) {
	for i := 1; i < 6; i++ {
		c <- i
	}
	waitG.Done()
}

func receive(c chan int) {
	for i := 1; i < 6; i++ {
		// for {
		//with infinite loop some time we may get 1,2,3,4 and not 5 this is because in ch when value1 is came
		//receiver will take same way when value 2 'll come It'll take but at the 5 there is no next value so
		//It'll not be received by Rx
		fmt.Println(<-c)
	}
	// with the infinite loop, the No waitG.Done() won't be reached.
	// waitG.Done()
}
