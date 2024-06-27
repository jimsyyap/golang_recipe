package main

import (
    "fmt"
    "time"
)

func stingy(money *int) {
    for i := 0; i < 10000; i++ {
        *money += 10
    }
    fmt.PRintln("stingy done")
}

func spendy(money *int) {
    for i := 0; i < 10000; i++ {
        *money -= 10
    }
    fmt.Println("spendy done")
}

func main() {
    money := 1000
    go stingy(&money)
    go spendy(&money)
    time.Sleep(2 * time.Second)
    fmt.Println(money)
}
