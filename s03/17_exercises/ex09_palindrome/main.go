// File name: ...\s03\exercises\set-8\palindrome
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

/*
Exercise: Write a program to read a name from concole
and determine whether it's a palindrome or not.

A palindrome word reads the same forward and backward.
Examples: level, madam, radar, mom, wow, nan

*/

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanf("%s ", &word)

	firstIdx := 0
	lastIdx := len(word) - 1

	isPalindrome := true
	for firstIdx < lastIdx {
		if word[firstIdx] != word[lastIdx] {
			isPalindrome = false
			break
		}
		firstIdx++
		lastIdx--
	}

	if isPalindrome {
		fmt.Printf("'%s' is a palindrome.\n", word)
	} else {
		fmt.Printf("'%s' is not a palindrome\n.", word)
	}

}
