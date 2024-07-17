package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printCount(label string, count chan int) {
	defer wg.Done()
	for {
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Println("Count: ", val, " received from ", label)
		if val == 10 {
			fmt.Println("Channel closed by ", label)
			close(count)
			return
		}
		val++
		count <- val
	}
}

func main() {
	count := make(chan int)
	wg.Add(2)

	go printCount("A", count)
	go printCount("B", count)
	fmt.Println("channel begin")
	count <- 1
	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
