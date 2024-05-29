package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("error read/write data...exit", err)
	}
}

/*
* this program is demonstrates reading and writing from standard input and writing to standard output using custom types in go. It prompts user to enter input, then echoes that input back to the screen with a prefix "out >". The io.copy function is used to handle the reading and writing.
 */
