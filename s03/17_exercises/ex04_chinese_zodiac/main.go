// File name: ...\s03\exercises\set-4\chinese_zodiac
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

/*
Exercise: Write a program to accept a year and
to determine the Chinese Zodiac sign for that year.

You may need to research this topic on the internet to
be able to write this code, but here are some hints:
- The Chinese Zodiac is based on a 12-year cycle;
- Each year represented by an animal, ...;
- Year 1900 is the year of the rat which is the
  fifth animal on the list.
*/

func main() {

	var year int

	fmt.Print("Enter a year: ")
	fmt.Scanf("%d ", &year)

	// List of animals: dog, pig, rat, ox, tiger,
	// rabbit, dragon, snake, horse, or sheep
	switch year % 12 {
	case 0: //examples: 1896, 1908, ...
		fmt.Println("monkey")
	case 1: //examples: 1897, 1909, ...
		fmt.Println("rooster")
	case 2: //examples: 1898, 1910, ...
		fmt.Println("dog")
	case 3: //examples: 1899, 1911, ...
		fmt.Println("pig")
	case 4: //examples: 1900, 1912, ...
		fmt.Println("rat")
	case 5: //examples: 1901, 1913, ...
		fmt.Println("ox")
	case 6: //examples: 1902, 1914, ...
		fmt.Println("tiger")
	case 7: //examples: 1903, 1915, ...
		fmt.Println("rabbit")
	case 8: //examples: 1904, 1916, ...
		fmt.Println("dragon")
	case 9: //examples: 1905, 1917, ...
		fmt.Println("snake")
	case 10: //examples: 1906, 1918, ...
		fmt.Println("horse")
	case 11: //examples: 1907, 1919, ...
		fmt.Println("sheep")
	}

}
