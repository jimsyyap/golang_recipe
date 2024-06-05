package main

import (
	"fmt"
	//"strconv"
)

func main() {

	counter := 0
pig:
	fmt.Println("Counter", counter)
	counter++
	if counter < 5 {
		goto pig
	}
}
