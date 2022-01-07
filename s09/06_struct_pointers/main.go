// File name: ...\s09\06_struct_pointers\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import "fmt"

type player struct {
	name, sport string
	age         int
}

func main() {

	p1 := &player{"Michael Phelps", "swimming", 32}
	fmt.Printf("(*p1).name=%s p1.name=%s\n", (*p1).name, p1.name)

	(*p1).age = 20
	fmt.Println("Player 1:", (*p1))

	player2 := player{"Diego Maradona", "Soccer", 57}
	p2 := &player2

	fmt.Printf("(*p2).name=%s p2.name=%s\n", (*p2).name, p2.name)
	fmt.Println("Player 2:", (*p2))

	test1()
}

func test1() {
	fmt.Println("\n\n")

	p1 := &player{"Michael Phelps", "swimming", 32}

	fmt.Printf("(*p1)=%v p1=%v\n", (*p1), p1)
	// output: (*p1)={Michael Phelps swimming 32} p1=&{Michael Phelps swimming 32}
	// (*p1) means the content of the address that p1 is pointing to.
	// p1 means the address (not the content of that address)

	fmt.Printf("(*p1).name=%s p1.name=%s\n", (*p1).name, p1.name)
	// output: (*p1).name=Michael Phelps p1.name=Michael Phelps
	// The first 'name' is an element of a content (so far ok).
	// The second one is an element of the content at that address
	// that is a bit differen than what you get from c pointers ...
	// The compiler somehow recognizes that a value has been requested...
	// automatic dereferencing ...

}
