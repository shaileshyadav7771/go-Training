// File name: ...\s11\03_concurrency_parallelism\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitG sync.WaitGroup
var cpuUsed = 1
var maxRandomNums = 100000
var counter = 0 // to check if race condition is happening or not?

var m sync.Mutex

//We can have init fun which will excecute before main() :)
//similar to constructor //we are using it for value initialization.

func init() {
	maxCPU := runtime.NumCPU() //It'll give us the max CPU :)

	cpuUsed = 1 //check by changing value's
	runtime.GOMAXPROCS(cpuUsed)

	fmt.Printf("Number of CPUs (Total=%d - Used=%d) \n", maxCPU, cpuUsed)
}

func main() {

	start := time.Now()
	ids := []string{"rotine1", "routine2", "routine3", "routine4"}

	waitG.Add(4)
	for i := range ids {
		go numbers(ids[i])
	}
	waitG.Wait()

	elapsed := time.Since(start)
	fmt.Printf("\nprogram took %s,%d. \n", elapsed, counter)
}

func numbers(id string) {
	rand.Seed(time.Now().UnixNano())

	m.Lock() //setting Mutex to check
	for i := 1; i <= maxRandomNums; i++ {
		counter++
		// time.Sleep(200 * time.Millisecond)
		fmt.Printf("%s-%d  ", id, rand.Intn(20)+20) // formula: rand.Intn(max-min) + min //range 40-20

	}
	m.Unlock()
	waitG.Done()
}
