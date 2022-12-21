package main

import (
	"fmt"
	"time"
)

/*
1. Find the issue with below code. Understand the root cause of it.
2. Fix the issue with using select `cancellation` or `timeout`.
*/
func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return completed
	}

	done := make(chan interface{})
	work := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-work
	fmt.Println("Done.")
}
