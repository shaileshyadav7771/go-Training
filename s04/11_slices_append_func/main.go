// File name: ...\s04\11_slices_append_func\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	team1 := []string{"Messi", "Pele", "Maradona", goalkeeper()}

	team := append([]string{"Maldini", "Cafu", "Carlos", "Beckenbauer"}, team1...)
	team = append(team, midfielders()...)

	for i, name := range team {
		fmt.Println(i, name)
	}

	fmt.Printf("Complete Slice is :%v,%T", team, team)
}

func midfielders() []string {
	return []string{"Dhoni", "Yadav", "Kohli"}
}

func goalkeeper() string {
	return "Shailesh"
}
