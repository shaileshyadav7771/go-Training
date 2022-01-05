// File name: ...\s02\exercises\set-3\future_day\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: If today is Sunday, what will be the day in 85 days?
// Sun=0, Mon=1, ..., Sat=6
//
// Examples:
// 	today=6, duration=8  future day=0
// 	today=6, duration=7  future day=6
// 	today=0, duration=7  future day=0
// 	today=0, duration=85 future day=1

func main() {

	today := 0     //Sun=0, Mon=1, ..., Sat=6
	duration := 85 //days
	futureDay := (today + duration) % 7

	fmt.Printf("today=%d, duration=%d future day=%d\n",
		today, duration, futureDay)
}
