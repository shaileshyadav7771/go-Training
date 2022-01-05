// File name: ...\s02\03_verbs\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {

	var k1 uint8 = 20

	f1 := 1234567.899
	fmt.Printf("%f - %T\n", f1, f1)
	fmt.Printf("%5.3f %.2f \n", f1, 214.437) // 	1234567.899 214.44

	//here 5 means 3.140(including . and %.2f means after . we need 2 digit)

	fmt.Printf("%5.3f %.2f \n", f1, 987654321.0987) // 1234567.899 987654321.10

	fmt.Println(k1+12, k1-12, k1*2, k1/2, 5%3)

	k1++
	fmt.Println(k1)
	k1 += 10
	fmt.Println(k1)

	fmt.Println(5/2, 5.0/2.0)

	// fmt.Println("\n\n")
	// fmt.Printf("%5.3f\n", 3.14)   //3.140
	// fmt.Printf("%10.8f\n", 3.14)  //3.14000000
	// fmt.Printf("%010.3f\n", 3.14) //000003.140
	// fmt.Printf("%10.3f\n", 3.14)  //bbbbb3.140   (b for a space character)
	// fmt.Printf("%-10.3f\n", 3.14) //3.140bbbbb   (b for a space character)
}
