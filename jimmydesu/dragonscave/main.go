/*
* You are in a land full of dragons. In front of you,
you see two caves. In one cave, the dragon is friendly
and will share his treasure with you. The other dragon
is greedy and hungry, and will eat you on sight.
Which cave will you go into? (1 or 2)
1
You approach the cave...
It is dark and spooky...
A large dragon jumps out in front of you! He opens his jaws and...
Gobbles you down in one bite!
Do you want to play again? (yes or no)
no
*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("You are in a land full of dragons. In front of you,")
	fmt.Println("you see two caves. In one cave, the dragon is friendly")
	fmt.Println("and will share his treasure with you. The other dragon")
	fmt.Println("is greedy and hungry, and will eat you on sight.")
	fmt.Println("Which cave will you go into? (1 or 2)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "1" {
		fmt.Println("You approach the cave...")
		fmt.Println("It is dark and spooky...")
		fmt.Println("A large dragon jumps out in front of you! He opens his jaws and...")
		fmt.Println("Gobbles you down in one bite!")
	} else if input == "2" {
		fmt.Println("You approach the cave...")
		fmt.Println("It is dark and spooky...")
		fmt.Println("A large dragon jumps out in front of you! He opens his jaws and...")
		fmt.Println("Gobbles you down in one bite!")
	}
}
