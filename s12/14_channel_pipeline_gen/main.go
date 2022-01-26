// File name: ...\s12\14_channel_pipeline_gen\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: A pipeline connects two (or more) channels operations, so that the
// result/output of the first channel operations becomes the input of the second channel.
//
// Write a program containing two functions to simulate a pipeline of two channels:
// step 1: function gen(numbers) sends some numbers to channel c.
// step 2: function sq(channel) reads those values from channel c and squares them.
// Ref: https://blog.golang.org/pipelines
//
func main() {
	nums := []int{2, 3, 4, 50000000999999}
	c := gen(nums...)
	fmt.Println("Printing c value is: ", c)
	out := sq(c)

	for range nums {
		fmt.Printf("%4d ", <-out)
	}

	fmt.Println()
	for n := range sq(sq(gen(nums...))) {
		fmt.Printf("%4d ", n)
	}
}

//receive only channel
func gen(nums ...int) <-chan int {
	channel1 := make(chan int) //bidirectional channel but return'll be receive only :)
	go func() {
		for _, n := range nums {
			channel1 <- n
		}
		close(channel1)
	}()

	return channel1
}

//Stage-2 pipeline
func sq(input <-chan int) <-chan int {
	channel2 := make(chan int) //bidirectional
	go func() {
		for n := range input {
			channel2 <- n * n
		}
		close(channel2)
	}()
	return channel2
}
