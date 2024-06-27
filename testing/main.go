package main

import (
    "fmt"
    "github.com/gonum/floats"
)

func main() {
    distance := floats.Distance([]float64{1, 2}, []float64{3, 4}, 2)
    fmt.Println(distance)
}
