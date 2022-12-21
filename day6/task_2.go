package main

import (
	"fmt"
	"time"
)

func main() {
	go random(5)
	time.Sleep(1 * time.Second)
	go random(10)
	time.Sleep(1 * time.Second)
	go random(15)
	time.Sleep(1 * time.Second)
}

func random(loop int) {
	go func() {
		for i := 0; i < loop; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Inside select statement")
		}
	}()
}
