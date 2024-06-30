package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

var key = []byte("some random key")

func checkMAC(message, recvMAAC []byte) bool {
    mac := hmac.New(sha256.New, key)
    mac.Write(message)
    calcMAC := mac.Sum(nil)

    return hmac.Equal(calcMAC, recvMAAC)
}

func main() {
    message := []byte("the red eagle flies at midnight")
    mac, _ := hex.DecodeString("f7b1b3b8b4f1b4b3b1b7")
    if checkMAC(message, mac) {
        fmt.Println("MAC verified")
    } else {
        fmt.Println("MAC verification failed")
    }
}
