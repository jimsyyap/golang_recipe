package main

import "fmt"

const (
	Info         = "\x1b[32m[INFO]\x1b[0m"
	Warn         = "\x1b[33m[WARN]\x1b[0m"
	Error        = "\x1b[31m[ERROR]\x1b[0m"
	Debug        = "\x1b[34m[DEBUG]\x1b[0m"
	Underline    = "\x1b[4m"
	UnderlineOff = "\x1b[24m"
	Italics      = "\x1b[3m"
	ItalicsOff   = "\x1b[23m"
)

func main() {
	fmt.Println(Info, "This is an info message")
	fmt.Println(Warn, "This is a warning message")
	fmt.Println(Error, "This is an error message")
	fmt.Println(Debug, "This is a debug message")
	fmt.Printf("%s%s%s\n\n", Italics, "Italics text", ItalicsOff)
	fmt.Printf("%s%s%s\n\n", Underline, "Underline text", UnderlineOff)
}
