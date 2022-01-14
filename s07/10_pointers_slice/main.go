package main

import "fmt"

// ASSIGNMENT - in previous sections we learned that a slice is a pointer to its
// underlying array. Create an array of strings and a slice of string.
// Using pointers prove that the slice uses the same address region of the array.
func main() {

	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("days=%v\n", days)
	fmt.Println("len of the day: ", len(days))
	fmt.Printf("Capacity of days:%d\n", cap(days))

	weekdays := days[4:7]
	fmt.Printf("the Value of weekdays(Slice) is: %v\n", weekdays)
	fmt.Printf("len(weekdays)=%d\n", len(weekdays))
	fmt.Printf("cap(weekdays)=%d\n\n", cap(weekdays))

	fmt.Printf("&days[5] Friday Address is: =%x\n", &days[4])
	fmt.Printf("&weekdays Friday Address at [0]=%x\n", &weekdays[0])

	//So the result for above both line is
	// &days[5] Friday Address is: =c0000b8040  // &weekdays Friday Address at [0]=c0000b8040
	//Which means both are pointing to same Address location(weekdays slice is part of days array).

	fmt.Println()
	fmt.Printf("&days[6] Sunday=%x\n", &days[6])
	fmt.Printf("&weekdays[2] sunday=%x\n", &weekdays[2])
}
