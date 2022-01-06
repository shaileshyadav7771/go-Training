// File name: ...\s03\exercises\set-6\random_cards1
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Exercise: Write a program to randomly pick a card
form a deck of 52 cards.
- rank: Ace, 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King
- suit: Clubs, Diamonds, Hearts, Spades

Sample output:
- You picked 5 of Clubs.
- You picked Queen of Hearts.
*/

func main() {

	const minRank = 1
	const maxRank = 13

	const minSuit = 1
	const maxSuit = 4

	rand.Seed(time.Now().UnixNano())
	rankIdx := minRank +
		rand.Intn(maxRank-minRank)
	suitIdx := minSuit +
		rand.Intn(maxSuit-minSuit)

	rankString := ""
	switch rankIdx {
	case 1:
		rankString = "Ace"
	case 2:
		rankString = "2"
	case 3:
		rankString = "3"
	case 4:
		rankString = "4"
	case 5:
		rankString = "5"
	case 6:
		rankString = "6"
	case 7:
		rankString = "7"
	case 8:
		rankString = "8"
	case 9:
		rankString = "9"
	case 10:
		rankString = "10"
	case 11:
		rankString = "Jack"
	case 12:
		rankString = "Queen"
	case 13:
		rankString = "King"
	}

	suitString := ""
	switch suitIdx {
	case 1:
		suitString = "Clubs"
	case 2:
		suitString = "Diamonds"
	case 3:
		suitString = "Hearts"
	case 4:
		suitString = "Spades"
	}

	fmt.Printf("You picked %s of %s.\n",
		rankString, suitString)
}
