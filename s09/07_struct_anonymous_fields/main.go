// File name: ...\s09\07_struct_anonymous_fields\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

// here we are only using Field Type instead of using It's name.
//becz of this we can't have one field of same type.
//This form is not recommonded but we are discussing for knowledge purpose.
type player struct {
	string
	int
}

func main() {

	player1 := player{"Shailesh Yadav", 77}
	fmt.Println("Player 1:", player1)

	//here we are accessing the value by using the variable Type.

	fmt.Printf("player1.int=%d player1.string=%s\n", player1.int, player1.string)

	player2 := player{ // Not to be used
		int:    36,
		string: "Federer",
	}
	fmt.Println("Player 2:", player2)

	// player3 := &player{"Munna", 07}
	// fmt.Println("Player 2:", player3)

}
