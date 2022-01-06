// File name: ...\s03\03_if_else_3\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	serverStatusOK := true

	if serverStatusOK {
		fmt.Println("Server is up & running.")
	}

	if s := "Production77"; serverStatusOK {
		fmt.Printf("%s server is up & running.", s)
		//another way of declaration is :
		// s := "Production77"
		//     if ; serverStatusOK {
		// 	     fmt.Printf("%s server is up & running.", s)

	} else {
		fmt.Printf("Sorry %s Server is not running. May be some problem", s)
	}

	// fmt.Println(s)	//compiler error: undefined
}
