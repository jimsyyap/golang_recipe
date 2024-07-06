package main

import (
    "fmt"
    "log"
    "os"
    "path"
)

func main() {
    LOGFILE := path.Join(os.TempDir(), "log.txt")
    fmt.Println(LOGFILE)
    f, err :=  os.OpenFile(LOGFILE, os.APPEND|os.CREATE|os.WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
        return
    }
    defer. f.Close()

    iLog := log.New(f, "customLog", log.LstdFlags)
    iLog.Println("This is a custom log message")
    ilog.Println("This is another custom log message")
}
