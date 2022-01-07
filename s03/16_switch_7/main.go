// File name: ...\s03\16_switch_7\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	var x interface{}

	// try other values for x: 10, 10.5, true, "hello", mult, no value, and uint(10)
	x = mult

	switch i := x.(type) {

	//in above case x.(type) means if we consider it to be 10 then that value'll be assign to
	//the i and then we are checking cases if i == int then we'll get matched case o/p
	case int:
		fmt.Printf("%T %v", i, i)
	case float64:
		fmt.Printf(" Value is for case float64: %T %v", i, i)
	case bool, string:
		fmt.Println("type is bool or string")
	case func(int) float64:
		fmt.Printf("value is for func(int):%T", i)
	case nil:
		fmt.Println("x is nil")
		//This will excecute if we don't define x variable value(we can check by commenting it.)
		//if we don't define any variable means zero value of interface() means nil.
	default:
		fmt.Println("don't know the type")
	}
}

//here we have defined a function and calling it from main loop and it return's float64
//so if pass int 10 then 10*1.5 = 15.0 so it will match with float case.
func mult(i int) float64 {
	return float64(i) * 1.5
}
