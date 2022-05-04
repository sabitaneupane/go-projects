package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Luck Number Game!
The program will pick %d random numbers.
Your mission is to guess one those numbers.

The greater your number is, the harder it gets.

Wanna play?	
`
)

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {

		n := rand.Intn(guess + 1)

		if n == guess {
			if turn == 1 {
				fmt.Println("FIRST TURN WINNER!!")
				return
			}

			switch rand.Intn(3) {
			case 0:
				fmt.Println("YOU WIN!!")
			case 1:
				fmt.Println("YOU'RE AWESOME!!")
			case 2:
				fmt.Println("CONGRATULATIONS!!")
			}
			return
		}
	}

	msg := "%s Try again?\n"
	switch rand.Intn(2) {
	case 0:
		fmt.Printf(msg, "☠️  YOU LOST...")
	case 1:
		fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
	}
}
