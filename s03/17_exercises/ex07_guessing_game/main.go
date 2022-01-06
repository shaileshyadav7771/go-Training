// File name: ...\s03\exercises\set-6\guessing_game
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Exercise: Write a program to generate a random number between n and m.
Then give the user x times to guess the number.
On each guess, let the user know if a bigger or smaller number
needs to be picked.
*/

func main() {
	const lowRange = 1
	const highRange = 10
	const numOfChances = 5

	rand.Seed(time.Now().UnixNano())
	randomNo := lowRange +
		rand.Intn(highRange-lowRange)

	// fmt.Println("** ", randomNo)

	fmt.Printf("You have %d chances to guess "+
		"a system-generated random number.\n", numOfChances)

	var yourNum int
	message := "You were not lucky this time. Try again!"
	for i := 0; i < numOfChances; i++ {
		fmt.Printf("Enter a number between %d and %d (try #%d): ",
			lowRange, highRange, i+1)
		fmt.Scanf("%d ", &yourNum)

		if yourNum == randomNo {
			message = "You guessed the number!"
			break
		} else {
			if yourNum > randomNo {
				fmt.Println("Pick a smaller number.")
			} else {
				fmt.Println("Pick a bigger number.")
			}
		}
	}

	fmt.Println("\n----------------------------------------------")
	fmt.Printf("-- Your Number=%d Random Number=%d \n-- %s",
		yourNum, randomNo, message)
	fmt.Println("\n----------------------------------------------")

}
