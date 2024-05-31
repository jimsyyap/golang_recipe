/*
* this is a simple greeting application that demonstrates how to parse command-line arguments, interact with the user to get input, and use that input to perform a task multiple times. It includes error handling to ensure the input is valid and to guide the user if something goes wrong.
 */

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes int
}

func GetName(r io.Reader, w io.Writer) (string, error) {
	scanner := bufio.NewScanner(r)
	msg := "What is your name?"
	fmt.Fprintf(w, msg)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("empty name")
	}
	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Hello, %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, "Hello, %s\n", name)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	name, err := GetName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}
	return nil
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.IntVar(&c.numTimes, "n", 0, "number of greetings")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() != 0 {
		return c, errors.New("positional arguments specified")
	}
	return c, nil
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
