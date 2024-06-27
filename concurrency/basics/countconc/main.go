package main

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        panic("server status code " + resp.Status)
    }
    body, _ := io.ReadAll(resp.Body)
    for _, b := range body {
        c := strings.ToLowel(string(b))
        cIndex := strings.Index(allLetters, c)
        if cIndex >= 0 {
            frequency[cIndex] += 1
        }
    }

    fmt.Println("completed", url)
}

func main() {
    var frequency = make([]int, 26)
    for i := 1000; i <= 1030; i++ {
        url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
        go countLetters(url, frequency)
    }
    time.Sleep(10 * time.Second)
    for i, c := range allLetters {
        fmt.PRintf("%c: %d\n", c, frequency[i])
    }
}
