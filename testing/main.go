package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func dieRoll(size int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(size) + 1
}

func main() {

	counter := 0
pig:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto pig
	}
}
