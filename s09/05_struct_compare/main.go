// File name: ...\s09\05_struct_compare\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"unsafe"
)

// ASSIGNEMENT: Comparing structs - considering the following struct write a
// simple program to compare two variables of this type. Write two versions
// of the compare() function: call by value and call by ref.

type book struct {
	name   string
	author string
}

func main() {
	b1 := book{"Go Programming", "Author 1"}
	b2 := book{"Go Programming", "Author 1"}
	b3 := book{"Python3 Programming", "Author 2"}

	fmt.Printf("LOC1 - %v\n", b1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	fmt.Println()

	fmt.Println(compare1(b1, b2))
	fmt.Println(compare1(b1, b3))
	fmt.Printf("LOC2 - %v\n\n", b1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	fmt.Println(compare2(&b1, &b2))
	fmt.Printf("LOC3 - %v\n\n", b1)

	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	fmt.Println()
	p1 := &book{"Go Web Programming", "Author 3"} //created pointer p1 and intialize

	//above p1 is same thing what we have defined below 2 line
	p2 := new(book)
	*p2 = book{"Go Web Programming", "Author 3"}

	fmt.Printf("LOC4 - &p1=%v *p1=%v\n", &p1, *p1)
	fmt.Println(compare2(p1, p2))
	fmt.Printf("LOC5 - &p1=%v *p1=%v\n", &p1, *p1)

	testSizeOfStruct()

}

func compare1(p book, q book) bool { // call by value
	if p.name == q.name && p.author == q.author {
		p.name = "Java Programming" // doesn't change the original variable
		return true
	}
	return false
}

func compare2(p *book, q *book) bool { // call by ref
	if p.name == q.name && p.author == q.author {
		p.name = "Scala Programming" // changes the original variable
		return true
	}
	return false
}

func testSizeOfStruct() {

	type values struct {
		val1 int32 // 4 bytes
		val2 int32 // 4 bytes
		val3 int64 // 8 bytes
		val4 int64 // 8 bytes
	}

	x := values{10, 20, 30, 40}
	p := &x
	fmt.Println(unsafe.Sizeof(p), unsafe.Sizeof(*p), unsafe.Sizeof(x)) //8 24 24
	//p is pointer and *p is value of that pointer and both value will be obeosily diferent.
	//here x is alreay there and then our p which pointing to it will cost more memory in low level programming
	//size of x and content of *p should be same.

	q := &values{11, 21, 31, 41}
	fmt.Println(unsafe.Sizeof(q), unsafe.Sizeof(*q)) //8 24

	//in first way we have declared the variable and then pointing to it(but due to additional memory cost for p) but in 2nd case
	//we are directly pointing to struct
}
