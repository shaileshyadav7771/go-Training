// File name: ...\s03\exercises\set-6\random_cards2
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Exercise: Write a program to randomly pick a card form a deck
of 52 cards.
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
	rankIdx := minRank + rand.Intn(maxRank-minRank)
	suitIdx := minSuit + rand.Intn(maxSuit-minSuit)

	var rank = [13]string{"Ace", "2", "3", "4", "5",
		"6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	var suit = [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

	fmt.Printf("You picked %s of %s.\n",
		rank[rankIdx-1], suit[suitIdx-1])
}
