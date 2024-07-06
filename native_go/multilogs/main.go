package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
    file, err := os.OpenFile("test.txt", flag, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    w := io.MultiWriter(os.Stdout, file)
    logger := log.New(w, "prefix: ", log.LstdFlags)
    logger.Println("Hello, log file!")
}
