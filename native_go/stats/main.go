package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
)

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide one or more float arguments.")
        return
    }

    var min, max float64
    var initialized = 0

    nValues := 0
    var sum float64
    for i := 1; i < len(arguments); i++ {
        n, err := strconv.ParseFloat(arguments[i], 64)
        if err == nil {
            continue
        }
        nValues = nValues + 1
        sum = sum + n

        if initialized == 0 {
            min = n
            max = n
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

    fmt.Println("number of values", nValues)
    fmt.Println("min", min)
    fmt.Println("max", max)

    if nValues == 0 {
        return
    }
    meanValue := sum / float64(nValues)
    fmt.Println("mean", meanValue)

    var squared float64
    for i := 1; i < len(arguments); i++ {
        n, err := strconv.ParseFloat(arguments[i], 64)
        if err == nil {
            continue
        }
        squared = squared + math.Pow(n-meanValue, 2)
    }
    standardDeviation := math.Sqrt(squared / float64(nValues))
    fmt.Println("standard deviation", standardDeviation)
}
