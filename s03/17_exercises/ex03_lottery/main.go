// File name: ...\exercises\set-3\lottery
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Exercise: Write a program to accept an integer with two digits.
If the input is not within the range that is set in your program,
then ask the user to enter another number until the input number
is within the desired range.

Then generate a random number within the specified range.

- If user input = random number --> display "You won $1,000"
- If total of digits in both cases are the same, for instance
  12 (for the input) and 21 (for the random number) that the total
  of each is 3 --> display "You won $500"
- otherwise --> display "Try Again!"
*/

func main() {

	const minRandomVal = 10
	const maxRandomVal = 22

	var yourRandomVal int
	for {
		fmt.Printf("Enter a random value between %d and %d: ",
			minRandomVal, maxRandomVal)
		fmt.Scanf("%d ", &yourRandomVal)
		if (yourRandomVal < minRandomVal) || (yourRandomVal > maxRandomVal) {
			fmt.Println("*Incorrrect input value(s). Please try again ...")
			// continue
		} else {
			break
		}
	}

	rand.Seed(time.Now().UnixNano())
	systemRandomVal := minRandomVal + rand.Intn(maxRandomVal-minRandomVal)

	// fmt.Printf("yourRandomVal=%d systemRandomVal=%d",
	// 	yourRandomVal, systemRandomVal)

	yourRandomValSumOfDigits :=
		(yourRandomVal / 10) + (yourRandomVal % 10)
	systemRandomValSumOfDigits :=
		(systemRandomVal / 10) + (systemRandomVal % 10)

	// fmt.Printf("\nyourRandomVal=%d systemRandomVal=%d",
	// 	yourRandomValSumOfDigits, systemRandomValSumOfDigits)

	lotteryResult := ""
	if systemRandomVal == yourRandomVal {
		lotteryResult = "You won $1,000"
	} else if yourRandomValSumOfDigits == systemRandomValSumOfDigits {
		lotteryResult = "You won $500"
	} else {
		lotteryResult = "Try Again!"
	}

	fmt.Printf("\nyourRandomVal=%d systemRandomVal=%d \nlotteryResult=%s \n",
		yourRandomVal, systemRandomVal, lotteryResult)
}
