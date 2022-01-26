// File name: ...\s12\15_channel_pipeline_words1\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	tjarratt "github.com/tjarratt/babble"
)

// ASSIGNMENT: Write a program including the following goroutines:
// 1) The 'newWords' goroutine that produces random words and send them to a channel;
// 2) The 'uWords' (uppercase words) goroutine that converts the randomly-generated
// words to uppercase;
// Use anonymous functions to write your goroutines.
//
func main() {
	start := time.Now()
	runtime.GOMAXPROCS(2) //testing

	babbler := tjarratt.NewBabbler() // requires "github.com/tjarratt/babble"
	new_word_generator_channel := make(chan string, 1000)
	Upperword_channel := make(chan string, 1000)

	go func() { // 'newWords' goroutine
		for i := 0; i < 50000; i++ {
			new_word_generator_channel <- babbler.Babble() //produces random words
		}

		close(new_word_generator_channel)
	}()

	go func() { // 'uWords' goroutine
		for w := range new_word_generator_channel {
			Upperword_channel <- w + " --> " + strings.ToUpper(w)
		}
		close(Upperword_channel)
	}()

	for w := range Upperword_channel {
		fmt.Println(w)

	}
	fmt.Printf("capacity:%d length:%d \n", cap(Upperword_channel), len(Upperword_channel))
	elapsed := time.Since(start)
	fmt.Printf("\nprogram took %s. \n", elapsed)
}
