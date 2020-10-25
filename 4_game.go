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
	fmt.Println(target)

	// Prompt the player to guess what the target is and store their response
	fmt.Println("Can you guess the number?")

	reader := bufio.NewReader(os.Stdin)
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
		fmt.Println("DONDAYO!!")
	}

}
