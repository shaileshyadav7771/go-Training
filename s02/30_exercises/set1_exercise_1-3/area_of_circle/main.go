// File name: ...\s02\exercises\set-1\area_of_circle\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

// Exercise: Write a program to calculate area of a circle.

// func main() {

// 	const PI = 3.14159
// 	var radius float64 = 3

// 	area := radius * radius * PI

// 	fmt.Printf("The area for the circle of radius %5.2f is %5.2f.\n",
// 		radius, area)
// }

//area of Circle is : pir2
func main() {
	pi := 3.14
	r := 7.0

	result := pi * r
	fmt.Printf("The area of Circle for redius %5.2f is %5.2f", r, result)

}
