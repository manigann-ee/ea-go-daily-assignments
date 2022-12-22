package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
}

func (numCounter *counter) Add(inpNumber int) {
	numCounter.count += inpNumber
}

// A simple program to trigger 10 goroutines
// & each to add 100 to Counter via iteration one by one
func main() {
	var numCounter counter
	group := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			//Instead of for loop we can just add 100 to the numCounter, not sure whether this is the issue
			//for i := 0; i < 100; i++ {
			//	numCounter.Add(1)
			//}
			numCounter.Add(100)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println(numCounter.count)
}
