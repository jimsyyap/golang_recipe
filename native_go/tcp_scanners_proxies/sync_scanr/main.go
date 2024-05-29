package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		wg.Add(1)
		go worker(ports, &wg)
	}
	for i := 0; i < 100; i++ {
		ports <- i
	}
	close(ports)
	wg.Wait()
}

/*
* this code creates a pool of 10 worker goroutiens
* it sends the numbers 0 to 99 to a channel
* each worker reads the numbers from the channel, prints them and signals completion
* the main function waits for all workers to finish before exiting
* in simple terms, this go code sets up to 100 workers to print numbers from 0 to 99 concurrently using channels to distribute the numbers and a waitgroup to wait for all the work to be completed
 */
