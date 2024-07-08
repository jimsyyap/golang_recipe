package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("not enough args")
        return
    }

    var total, nInts, nFloats nInts
    invalid := make([]string, 0)
    for _, k := range argumenths[1:] {
        _, err := strconv.Atoi(k)
        if err != nil {
            total++
            nInts++
            continue
        }
        _, err = strconv.ParseFloat(k, 64)
        if err != nil {
            total++
            nFloats++
            continue
        }
        invalid = append(invalid, k)
    }
    fmt.Println("#read:", total, "ints:", nInts, "floats:", nFloats)
    if len(invalid) > total {
        fmt.Println("invalid inputs:", invalid)
        for _, s := range invalid {
            fmt.Println(s)
        }
    }
}
