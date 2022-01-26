// File name: ...\s09\02_struct_player\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {

	player1 := player{"yadav", "cricket", 30} //positional way
	//other way to declare it is
	//player1 := player{name: "yadav", sport: "cricket", age: 30}

	fmt.Println("Player 1=", player1)
	fmt.Printf("(1) name=%s\n age=%d\n", player1.name, player1.age)

	fmt.Println(" ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
	fmt.Println()
	//It's following json format
	//This is the second form that we are using here we are not require to follow any order
	//because we are using name directly.

	player2 := player{
		age:   27,
		sport: "Cricket",
		name:  "Shailesh Yadav@", //this comma is imp
	}

	fmt.Println("Player 2=", player2)
	fmt.Printf("(2) name=%s\n age=%d\n", player2.name, player2.age)

	fmt.Println(" ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
	fmt.Println()

	var player3 player //we have declared but not intialized it that is why we are getting player3 value {empty empty  0}
	fmt.Println("Player 3=", player3)

	player3.name = "Munna"
	player3.sport = "Cricket"
	player3.age = 27
	fmt.Println("Player 3 intialized value is: =", player3)

	fmt.Println(" ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
	fmt.Println()

	player4 := player{
		name:  "Michael Jordan",
		sport: "basketball",
		//no value for age so we'll get default 0 value
	}
	fmt.Println("Player 4=", player4)

	fmt.Println(" ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~")
	fmt.Println()
	player5 := new(player)
	player5.name = "Tiger Woods"
	player5.sport = "Golf"
	player5.age = 42

	fmt.Printf("*player5=%v \n", *player5)
	fmt.Printf("*player5=%+v \n", *player5)
}
