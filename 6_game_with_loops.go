package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Generate a random number from 1 to 100 and store  as target
	rand.Seed(time.Now().UTC().UnixNano())
	target := rand.Intn(100) + 1

	// Prompt the player to guess what the target is and store their response
	fmt.Println("Can you guess the number?")
	fmt.Println("Choose a number between 1 and 100")

	reader := bufio.NewReader(os.Stdin)
	guesses := 10
	success := false

	for i := guesses; i >=1 ; i-- {
		fmt.Println("You have", i, "guesses left")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		// If the player guess is less than the target number say "Oops, Your guess was LOW"
		// If the player guess is greater than the target number say "Oops, Your guess was HIGH"
		if guess < target {
			fmt.Println("Oops, Your guess was LOW")
		} else if guess > target {
			fmt.Println("Oops, Your guess was HIGH")
		} else {
			success = true
			fmt.Println("YOU NAILED IT!!")
			break
		}
	}

	if !success {
		fmt.Println("The number was", target)
	}

}
