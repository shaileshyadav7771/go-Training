// File name: ...\s03\exercises\set-2\leap_year
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

/*
Exercise: Write a program to allow the user to enter a
year and then to determine whether it's a leap year or not.

There are 30 leap years between 1900 and 2020:
1904, 1908, 1912, 1916, 1920, 1924, 1928, 1932,
1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964,
1968, 1972, 1976, 1980, 1984, 1988, 1992, 1996,
2000, 2004, 2008, 2012, 2016, 2020

*/

func main() {

	var year int
	fmt.Print("Enter a year:")
	fmt.Scanf("%d ", &year)

	isLeapYear :=
		(year%4 == 0 && year%100 != 0) || (year%400 == 0)

	if isLeapYear {
		fmt.Printf("\nYear '%d' is a leap year.\n", year)
	} else {
		fmt.Printf("\nYear '%d' is not a leap year.\n", year)
	}
}
