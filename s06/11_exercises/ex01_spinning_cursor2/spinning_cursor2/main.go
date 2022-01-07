// File name: ...\s03\exercises\set-5\spinning_cursor2
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

/*
Exercise: Re-write the 'spinning cursor' program using functions.
This time, instead of stopping the program using CTRL + C, let
the program runs for 'n' seconds and then exits automatically.
Report the elapsed time as well.
*/

// Choose a smaller numer to make the spinnig faster.
var delayTime = 100 * time.Millisecond

func main() {
	executionTime := 3 * time.Second
	start := time.Now()

	fmt.Printf("Program will end in about %v.\n",
		executionTime)
	fmt.Print("Operation in progress ... ")

	s := `\|/-`
	i := 0
	for {
		printSpinningSymbol(string(s[i]))

		if time.Since(start) > executionTime {
			fmt.Println("Done")
			fmt.Println("Elapsed Time: ",
				time.Since(start))
			break
		}

		i = (i + 1) % 4
	}
}

func printSpinningSymbol(symbol string) {
	fmt.Print(symbol)
	time.Sleep(delayTime)
	fmt.Print("\b")
}
