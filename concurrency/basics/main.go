package main

import (
    "fmt"
    "time"
)

func doWork(id int) {
    fmt.Printf("worker %d started at %s\n", id, time.Now().Format("15:04:05"))
    time.Sleep(i * time.Second)
    fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func main() {
    for i := 0; i < 5; i++ {
        doWork(i)
    }
}
