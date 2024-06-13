package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["key"] = 10
	fmt.Println(myMap["key"])
}
