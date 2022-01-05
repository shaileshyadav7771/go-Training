package main

import "fmt"

func main() {

	fmt.Printf("%d,%o,%x,%b", byte('A'), byte('A'), byte('A'), byte('A'))

	//Now another way to represent the same is by using [] bracket means %d will take first byte('A')
	//And [1] here one means consider the first vale which is byte('A')

	fmt.Printf("%d %[1]o %[1]x %[1]b \n", byte('A'))

	fmt.Printf("%d %o %x %b \n", 'A', 'A', 'A', 'A')

}

// o/p:
// PS C:\Users\shailesh_yadav\go Training\05_nosystem> go run .\main.go
// 65,101,41,100000165 101 41 1000001
// 65 101 41 1000001
