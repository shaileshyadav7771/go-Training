// File name: ...\s04\01_arrays_int\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
)

func main() {
	var nums [3]int

	var sum1 int
	var sum2 int

	fmt.Printf("Array is having value:%v , Type=%T  and len=%d\n", nums, nums, len(nums))

	for i := range nums {
		sum1 += i //in go if we have defined any variable and value is not assign to it then it's value 'll be
		//0 that is why sum1 += i is sum1 = sum1 +i i.e 0+0(i value)
		sum2 += nums[i] // here nums[i] means nums[0] so nums[0] >It's value is zero in array
	}

	fmt.Println(sum1, sum2)
}
