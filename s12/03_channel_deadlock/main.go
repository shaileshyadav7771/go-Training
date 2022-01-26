// File name: ...\s12\03_channel_deadlock\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

func main() {
	c := make(chan string)
	c <- "No one likes my channel!" //fatal error: all goroutines are asleep - deadlock!

	//here no one is recieve details from the channel
	//A key note is that a programme go into deadlock because of bad design like above
}
