// File name: ...\s09\13_struct_method-overload\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"strings"
)

type movie struct {
	name  string
	actor string
}

type imdb struct {
	movie
	name    string
	comment string
}

//In go we can not use same fn name (overloading is not allowed).
// func display() {
// 	fmt.Println("Hello")
// }
// func display(msg string) {
// 	fmt.Println(msg)
// }

//in below function we can see funct fullInfo() with same name
//but both are having different receiver and in Go programming Receiver is also a part of signature that is why
//we are not getting any ERROR.

func (m movie) fullInfo() string {
	return m.name + "-" + m.actor
}

func (i imdb) fullInfo() string {
	return strings.ToLower(i.movie.name + "-" +
		i.movie.actor + "-" + i.name + "-" + i.comment)
}

//Note above type is called overriding becz both are associted with different strct (Like how we see in OOP'S)
//but in some books they are calling it as method overloading and at some place overriding. :)

func main() {
	m := movie{"Bahubali", "Shailesh"}

	i1 := imdb{m, "Sony Pictures", "Loved it!"}

	//another way to declare is:
	i2 := imdb{
		movie: movie{
			name:  "The Godfather",
			actor: "Marlon Brando",
		},
		name:    "Paramount Pictures",
		comment: "Superb!",
	}

	fmt.Println(m.fullInfo())
	fmt.Println(i1.fullInfo())
	fmt.Println(i2.fullInfo())

	fmt.Printf("%s-%s \n", i2.movie.name, i2.name)

	p1 := &movie{"Forrest Gump", "Tom Hanks"}
	fmt.Printf("%T - %v - name=%s - actor=%s\n", p1, *p1, p1.name, p1.actor)

	p2 := &i2
	fmt.Printf("%T - %v\n", p2, *p2)

	test1()
}
func test1() {
	fmt.Println("\n\n")
	// 'm' is a memory address that has a name
	m := movie{"x1", "x2"}

	// 'p1' is a pointer to a memory address that has a name.
	// This means that you have access to that memory address
	// using its name and also using the pointer 'p1.
	p1 := &m
	fmt.Printf("%T - %v - name=%s - actor=%s\n", p1, *p1, p1.name, p1.actor)
	// output: *main.movie - {x1 x2} - name=x1 - actor=x2

	// 'p2' is a pointer to a memory address that doesn't have a name.
	// This is called a dangling memory address in some languages.
	// After the following assignment, you can only access this
	// memory address using the pointer.
	p2 := &movie{"y1", "y2"}
	fmt.Printf("%T - %v - name=%s - actor=%s\n", p2, *p2, p2.name, p2.actor)
	// output: *main.movie - {y1 y2} - name=y1 - actor=y2
}
