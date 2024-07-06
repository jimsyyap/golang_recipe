package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: main <number>")
        return
    }
    argument := os.Args[1]

    switch argument {
    case "0":
        fmt.Println("Zero")
    case "1":
        fmt.Println("One")
    case "2", "3", "4":
        fmt.Println("Two, Three or Four")
        fallthrough
    default:
        fmt.Println("Unknown number: " + argument)
    }

    value, err := strconov.Atoi(argument)
    if err != nil {
        fmt.Println("Error: " + err.Error())
        return
    }

    switch {
    case value == 0:
        fmt.Println("Zero")
    case value > 0:
        fmt.Println("Positive number")
    case value < 0:
        fmt.Println("Negative number")
    default:
        fmt.Println("Unknown number: ", value)
}
