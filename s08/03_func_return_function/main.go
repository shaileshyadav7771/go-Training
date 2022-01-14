// File name: ...\s08\03_func_return_function\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

func main() {
	hi := sayGreetings("Marathi")
	fmt.Println(hi())

	hi = sayGreetings("GER")
	//above we have declared but not called our fn :). so below line we are calling directly by passing value ()
	fmt.Println(sayGreetings("Bhojpuri")())

	fmt.Println()
	fmt.Printf("%T \n", hi)   //Note: Return Type of this Fn is > func() string
	fmt.Printf("%T \n", hi()) //Return Type of this fn is > string
}

//below is sayGreetings is function which is returning another function func()> return string(func() string in first line)
func sayGreetings(lang string) func() string {

	hi := func() string {
		return "Hello- English"
	}
	//above fn is the default fn which will excecute automatically and return string "Hello"

	if lang == "Marathi" {
		hi = func() string {
			return "Kay bhava kasa Aahet.."
		}
	} else if lang == "Bhojpuri" {
		hi = func() string {
			return "Ka haal Ba"
		}
	}

	return hi
}
