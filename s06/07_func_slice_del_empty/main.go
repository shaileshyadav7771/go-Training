// File name: ...\s06\07_func_slice_del_empty\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// ASSIGNMENT: Write a function to accept a slice of
// string, and return a new slice with no blank strings.
// Use both variadic and regular functions.

//Imp : we can acheive this by both way's :).
func main() {

	data := []string{"Daisy", "Rose", "", "Tulip"}

	fmt.Printf("NewData ki value after return %q\n", trimSlice(data)) //form 1
	// fmt.Printf("%q\n", trimSlice(data...)) //form 2

	fmt.Printf("Original data value after all change%q\n", data)
	// fmt.Println("Printing data without %q is :", data) //[Daisy Rose  Tulip]
}

func trimSlice(data1 []string) []string { //form1
	// func trimSlice(data ...string) []string { //form2
	data1[2] = "Shailesh"
	fmt.Println(data1)
	newData := data1[:0]

	fmt.Println("newData ki value is : ", newData)
	//Note above is different It will change value in Main data string(refernce Type)
	// var newData []string

	for _, d := range data1 {

		if d != "" {
			newData = append(newData, d)
		}
	}
	return newData
}
