package main

import (
    "fmt"
    "sync"
    "time"
)

func doWork(id int, wg *sync.WaitGroup, ch chan string) {
    defer wg.Done() //signal completion to the waitgroup

    fmt.Printf("Work %d started at %s\n", id, time.Now().Format("04:05"))
    time.Sleep(2 * time.Second)
    fmt.Printf("Work %d completed at %s\n", id, time.Now().Format("04:05"))

    ch <- fmt.Sprintf("Work %d completed", id) //send a message to the channel
}

func main() {
    const numWorkers = 15
    var wg sync.WaitGroup
    wg.Add(numWorkers) //wait for all the workers to complete

    ch := make(chan string)

    for i := 0; i < numWorkers; i++ {
        go doWork(i, &wg, ch)
    }

    // receive completion msg from the channel
    for i := 0; i < numWorkers; i++ {
        message := <-ch
        fmt.Println(message)
    }

    wg.Wait()
    fmt.Println("All work completed")
}
