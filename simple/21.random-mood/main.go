package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var mood = []string{
	"happy 😀",
	"good 👍",
	"awesome 😎",
	"sad 😞",
	"bad 👎",
	"terrible 😩",
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please input your name.")
		return
	}

	name := os.Args[1]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(mood))

	log.Printf("%s is %s.", name, mood[n])
}
