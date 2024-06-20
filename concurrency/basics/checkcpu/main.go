package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Go version: " + runtime.Version())
    fmt.Println("Go OS: " + runtime.GOOS)
    fmt.Println("Go Architecture: " + runtime.GOARCH)
    fmt.Println("Go Root: " + runtime.GOROOT())
    fmt.Println("Go NumCPU: ", runtime.NumCPU())
    fmt.Println("Go NumGoroutine: ", runtime.NumGoroutine())
    fmt.Println("Go NumCgoCall: ", runtime.NumCgoCall())
    // fmt.Println("Go NumGc: ", runtime.NumGc())
    fmt.Println("Go MemStats: ", runtime.MemStats{})
    fmt.Println("Go Compiler: ", runtime.Compiler)
    fmt.Println("Go GOMAXPROCS: ", runtime.GOMAXPROCS(0))
}
