package main

import (
    "fmt"
    "log"
    "os"
    "path"
)

func main() {
    LOGFILE := path.Join(os.Getenv("HOME"), "log.txt")
    fmt.Println("Logging to", LOGFILE)
    f, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("error opening file: %v", err)
    }
    defer f.Close()

    LstdFlags := log.Ldate | log.Ltime | log.Lshortfile
    iLog := log.New(f, "INFO: ", LstdFlags)
    iLog.Println("This is an info log message")
    iLog.SetFlags(log.Lshortfile | log.LstdFlags)
    iLog.Println("This is an info log message with short file name")
}
