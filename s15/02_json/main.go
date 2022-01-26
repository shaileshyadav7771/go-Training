// File name: ...\s15\02_json\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// SoccerClubs ...
type SoccerClubs struct {
	Name    string
	Country string
	Value   float32 `json:"Dollar Value(B)"` //It's alias name which we are telling to compiler and we will see in JSON
	Players []string
}

var teams = []SoccerClubs{
	{Name: "Manchester United", Country: "England", Value: 3.69,
		Players: []string{"Eric Cantona", "David Beckham", "Ryan Giggs"}},
	{Name: "Barcelona", Country: "Spain", Value: 3.64,
		Players: []string{"Johan Cruyff", "Xavi", "Andrés Iniesta"}},
	{Name: "Real Madrid", Country: "Spain", Value: 3.58,
		Players: []string{"Zidane", "Roberto Carlos", "Ronaldo"}},
}

func main() {
	// fmt.Println("teams:", teams)
	fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - ")

	// use either json.Marshal() or json.MarshalIndent()
	// data, err := json.Marshal(teams)
	//Marshaling means converting from go lang data structure form to json Data form.
	//so data is json form
	data, err := json.MarshalIndent(teams, "", "   ")

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%v\n\n", teams)
	fmt.Printf("%+v\n\n", teams) //+v means It will show key name also.
	fmt.Printf("JSON Data is: %s\n", data)

	fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - ")
	var names []struct {
		Name    string
		Country string
	}

	//Now we are converting from JSON Data type to Go structure.
	if err := json.Unmarshal(data, &names); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(names)

	var a string = "s"
	fmt.Println(a)

}
