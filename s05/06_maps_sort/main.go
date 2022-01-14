// File name: ...\s05\06_maps_sort\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"sort"
)

// ASSIGNEMENT: Maps are unordered. Write a program to use
// a "map[string]float64" data map and sort it.
func main() {
	sal := map[string]float64{
		"Blake":  60000.00,
		"Parker": 120000.50,
		"Dakota": 93000.00,
	}
	// the order is not guranteed for below excecution.
	fmt.Println("The Dict is: 1 :", sal)
	// fmt.Println("The Dict is :", sal)
	// fmt.Println("The Dict is :", sal)
	// fmt.Println("The Dict is : 4", sal)

	names := make([]string, 0, len(sal))
	fmt.Printf(" names of value: %T\n", names)

	for n := range sal {
		// fmt.Println("Testing:", n, s)
		names = append(names, n)
	}
	fmt.Println(names)

	sort.Strings(names)
	fmt.Println(names)

	for _, n := range names {
		fmt.Printf("%s\t%.2f\n", n, sal[n])
	}

}
