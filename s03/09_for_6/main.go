// File name: ...\s03\09_for_6\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Printf("\ni=%d ** ", i)

		for j := 1; j < 3; j++ { //This loop will excecte 2 times for i=1 (j=1,2) :) trace on paper .
			fmt.Printf("%d ", i*j)
		}
	}
}
