package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader implements io.Reader
type FooReader struct{}

// Read reads from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out > ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("read %d bytes from stdin\n", s)

	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("wrote %d bytes to stdout\n", s)
}

/*
* FooReader - custom reader that reads from the standard input and prompts "in >"
* FooWriter - custom writer that writes to the standard output and prompts "out >"
* main - reads from stdin and writes to stdout
* this program is a simple demo of how to implement custom I/O operations is go using interfaces and struct methods
 */
