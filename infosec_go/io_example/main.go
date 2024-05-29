package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (f FooReader) Read(p []byte) (n int, err error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter FooWriter) Write(p []byte) (n int, err error) {
	fmt.Print("out > ")
	return os.Stdout.Write(p)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)

	s, err := reader.Read(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes\n", s)

	s, err = writer.Write(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes\n", s)
}

/*
*This program reads input from the user, displays the input length, and then writes the same input back to the screen, again displaying the number of bytes written. It's a simple demonstration of custom reading and writing in Go.
 */
