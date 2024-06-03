package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("file.txt", os.O_RDONLY, 0666)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err != nil {
			return
		}
	}
}
