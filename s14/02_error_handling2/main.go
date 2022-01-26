// File name: ...\s14\02_error_handling2\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	_, _, err := produceError1(true)
	if err != nil {
		fmt.Println("test1")
		// log.Fatalln("err ki value is: ", err)
	} else {
		log.Println("No error!")
	}

	_, _, error := produceError2(true)
	if error != nil {
		fmt.Println("test2")
		// log.Fatalln("error ki value is: ", error)
	} else {
		fmt.Println("******************************")
		log.Println("Sorry No ERROR found :) ")

	}

}

func produceError1(b bool) (int, int, error) {

	if !b {
		return 0, 0, nil
	}
	fmt.Println()
	errMsg1 := errors.New("ProduceError1- b value is false so excecuting this line.. ")
	fmt.Printf("-errors.New() generates type: %T\n", errMsg1)

	return 0, 0, errMsg1
}

func produceError2(b bool) (int, int, error) {

	if !b {
		return 0, 0, nil
	}
	fmt.Println()
	error := fmt.Errorf("it's %v that I'm just an error value", b)
	fmt.Printf("**errors.New() generates type: %T\n", error)

	return 0, 0, error
}
