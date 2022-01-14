// File name: ...\s06\03_func_slice\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	scores1 := []float32{91, 82, 99} // slice is of refernce type
	fmt.Println("score1 value is: ", scores1)
	scores1[1] = 7958
	fmt.Println("score1 value after changing 1 th index to 7958 is: ", scores1)
	fmt.Printf("average: %.2f,%T\n", avg(scores1), scores1)
	//Now let change any one value after fn call(inside fn) and see it's value as It's refernce type.
	fmt.Println("score1 value after Fn call is: ", scores1)

	scores2 := []float32{72, 81, 78, 91, 68}
	fmt.Printf("average: %.5f\n", avg(scores2))

	println("********** Testing ***********")
	var s string
	s = "abcd"
	fmt.Print(s)
	s[0] = 99999
	fmt.Print(s)
}

//Note: We are passing value with score1 name and in our function it's score so name can be different.
func avg(scores []float32) float32 {
	var total float32
	for _, score := range scores {
		total += score
	}
	scores[0] = 7771
	return total / float32(len(scores))
}

// In Golang, one of the confusing part is value type and reference type.

// Value type: int, float, string, bool, structs

// Reference type: slides, maps, channels, pointers, functions
