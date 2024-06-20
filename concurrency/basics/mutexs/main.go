package main

import (
    "fmt"
    "sync"
    "time"
)

type SharedData struct {
    counter int
    mu sync.Mutex
}

func (sd *SharedData) increment() {
    sd.mu.Lock()
    defer sd.mu.Unlock()
    sd.counter++
}

func doWorkWithMutex(id int, data *SharedData, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("work %d started at %s\n", id, time.Now().Format("15:04:05"))
    time.Sleep(2 * time.Second)

    data.increment()

    fmt.Printf("work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
    const numWorkers = 15
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    sharedData := &SharedData{}

    for i := 0; i < numWorkers; i++ {
        go doWorkWithMutex(i, sharedData, &wg)
    }

    wg.Wait()

    fmt.Printf("counter: %d\n", sharedData.counter)
}
