// File name: ...\s02\exercises\set-3\show_current_time\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

//
// Exercise: In Go you can get Unix time in milliseconds using the following:
// time.Now().UnixNano() / int64(time.Millisecond)		("time" package)
// Write a program to get Unix time in milliseconds and generate the following output:
//
// 	totalMilliSeconds=1544851401366
// 	totalSeconds=1544851401
// 	currentSecond=21
// 	totalMinutes=25747523
// 	currentMinute=23
// 	totalHours=429125
// 	currentHour=5
//
// You can use the following site that displays Milliseconds since Epoch.
// https://currentmillis.com/
// Compare the result of the program with the value on the site.

func main() {

	// Unix time in Go in milliseconds - Milliseconds since Epoch
	// total seconds since midnight, Jan 1, 1970
	totalMilliSeconds :=
		time.Now().UnixNano() / int64(time.Millisecond)

	totalSeconds := totalMilliSeconds / 1000

	currentSecond := totalSeconds % 60

	totalMinutes := totalSeconds / 60

	currentMinute := totalMinutes % 60

	totalHours := totalMinutes / 60

	currentHour := totalHours % 24

	fmt.Printf("totalMilliSeconds=%d \n", totalMilliSeconds)
	fmt.Printf("totalSeconds=%d \n", totalSeconds)
	fmt.Printf("currentSecond=%d \n", currentSecond)
	fmt.Printf("totalMinutes=%d \n", totalMinutes)
	fmt.Printf("currentMinute=%d \n", currentMinute)
	fmt.Printf("totalHours=%d \n", totalHours)
	fmt.Printf("currentHour=%d \n", currentHour)
}
