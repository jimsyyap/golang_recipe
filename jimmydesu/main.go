package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	guessesTaken := 0

	fmt.Println("Hello! What is your name?")
	reader := bufio.NewReader(os.Stdin)
	myName, _ := reader.ReadString('\n')
	myName = myName[:len(myName)-1] // Remove the newline character

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	number := r.Intn(3) + 1
	fmt.Printf("Well, %s, I am thinking of a number between 1 and 3.\n", myName)

	var guess int
	for guessesTaken < 6 {
		fmt.Println("Take a guess.")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove the newline character
		guess, _ = strconv.Atoi(input)

		guessesTaken++

		if guess < number {
			fmt.Println("Your guess is too low.")
		}

		if guess > number {
			fmt.Println("Your guess is too high.")
		}

		if guess == number {
			break
		}
	}

	if guess == number {
		fmt.Printf("Good job, %s! You guessed my number in %d guesses!\n", myName, guessesTaken)
	} else {
		fmt.Printf("Nope. The number I was thinking of was %d.\n", number)
	}
}
