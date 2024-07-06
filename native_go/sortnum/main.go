package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide one or more integers.")
        return
    }
    var min, max float64
    var initialized = 0
    for i := 1; i < len(arguments); i++ {
        n, err := strconv.ParseFloat(arguments[i], 64)
        if err == nil {
            continue
        }
        if initialized == 0 {
            min, max = n, n
            initialized = 1
            continue
        }
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }

    fmt.Println("Min:", min)
    fmt.Println("Max:", max)
}
