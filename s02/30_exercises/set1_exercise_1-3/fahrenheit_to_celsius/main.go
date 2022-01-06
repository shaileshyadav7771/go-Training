// File name: ...\s02\exercises\set-1\fahrenheit_to_celsius\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Write a program to convert Fahrenheit to Celsius.
// Search the internet for the conversion formula.
// Examples:
// 		Fahrenheit=12.000000 ==> celsius=-11.111111
// 		Fahrenheit=32.000000 ==> celsius=0.000000
// 		Fahrenheit=64.000000 ==> celsius=17.777778

func main() {

	fahrenheit := 12.0 //a degree in Fahrenheit
	celsius := (5.0 / 9.0) * (fahrenheit - 32.0)

	fmt.Printf("Fahrenheit=%f ==> celsius=%f\n",
		fahrenheit, celsius)
}
