/*
Student's question:
Should the playerInfo function signature not be like this

    type athletes interface{}

    // func playerInfo(a interface{}) {
    //  fmt.Println(a)
    // }

    func playerInfo(a athletes) {
        fmt.Println(a)
	}


Answer:
Thank you for your question.
The result will be the same in both scenarios.
This is because 'athletes' is of the same type (empty interface).
Saying that, using 'a athletes' is more readable, and it was the reason that I’ve originally wrote the 'type' line, right before the 'playerInfo' function section.
In addition to readability, it's also closer to the syntax of other programming languages.

To provide with more examples, I have added a small snippet to the section’s program, tagged by 'Added on 15-Oct-2018'.


*/

// File name: ...\s10\02_interface_empty\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

type athlete struct {
	name    string
	country string
}

type football struct { //meaning footbalPlayer!
	athlete
	position string
}

type tennis struct {
	athlete
	rightHanded bool
}

type athletes interface{}

// func playerInfo(a interface{}) {
func playerInfo(a athletes) {
	fmt.Println(a)
}

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// Added on 15-Oct-2018
type numType struct {
	val interface{}
}

// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~

func main() {
	messi := football{}
	pele := football{}
	federer := tennis{}
	nadal := tennis{}

	favAthletes := []athletes{messi, pele, federer, nadal}

	for k, v := range favAthletes {
		fmt.Println(k, " - ", v)
	}

	messi = football{athlete{"Leo Messi", "Argentina"}, "Attcker"}
	federer = tennis{athlete{"Roger Federer", "Switzerland"}, true}
	playerInfo(messi)
	playerInfo(federer)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	pele = football{athlete{"Pele", "Brazil"}, "Attcker"}
	nadal = tennis{athlete{"Rafael Nadal", "Spain"}, false}

	favAthletes2 := []interface{}{messi, pele, federer, nadal}

	fmt.Println(favAthletes2)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	// Added on 15-Oct-2018
	fmt.Println()
	nt := numType{}

	nt.val = "two"
	fmt.Printf("%#v %T \n", nt.val, nt.val) // "two" string

	nt.val = 2
	fmt.Printf("%#v %T \n", nt.val, nt.val) //2 int

	nt.val = messi
	fmt.Printf("%#v %T \n", nt.val, nt.val)
	// main.football{athlete:main.athlete{name:"Leo Messi", country:"Argentina"}, position:"Attcker"} main.football
}
