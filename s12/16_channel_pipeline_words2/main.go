// File name: ...\s12\16_channel_pipeline_words2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strings"

	tjarratt "github.com/tjarratt/babble"
)

// ASSIGNMENT: Rewrite the previous assignment (pipeline-like)
// using functions and unidirectional Channels.
func main() {

	new_word_channel := make(chan string)
	uppercase_channel := make(chan string) //uppercase words

	go sendWords(new_word_channel)
	go convertWords(uppercase_channel, new_word_channel)
	//printing upper channel Data:

	printWords(uppercase_channel)
}

func sendWords(out chan<- string) {
	babbler := tjarratt.NewBabbler()
	for i := 0; i < 5; i++ {
		out <- babbler.Babble() //produces random words
	}
	close(out)
}

func convertWords(out chan<- string, in <-chan string) {
	for w := range in {
		out <- w + " --> " + strings.ToUpper(w)
	}
	close(out)
}

func printWords(in <-chan string) {
	for w := range in {
		fmt.Println(w)
	}
}
