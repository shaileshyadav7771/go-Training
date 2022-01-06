// line comment

/*
Multi
Line
comment

*/

package main

import "fmt"

func main() {
	fmt.Println("Hello This is our first Go pragramme as part of Go Training.₹.")
	concept()

}

func concept() {
	firstName := "Shailesh"
	firstName = "10\t"
	case1 := "Yadav\a" //not recommonded as per the go lang.
	fmt.Println(firstName, case1)

	// diff between "" ' '
	firstName = "Shailu\n\n" //two blank line'll be added after printiong Shailu
	// firstName = `Munna\n\n`  //  o/p : Munna\n\n
	fmt.Println(firstName)

	a := 11.0
	fmt.Printf("%d,%x,%T", a, a, a)
}
