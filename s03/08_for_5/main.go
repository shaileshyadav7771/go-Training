// File name: ...\s03\08_for_5\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Press CTRL+C to stop the program.")

	i := 2
	for {
		time.Sleep(2 * time.Second) //here 2 is the duration which means 2 sec :).
		fmt.Println(i, " ")
		i++
	}
}
