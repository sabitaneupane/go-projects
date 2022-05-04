package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var moods = [2][3]string{
	{"happy 😀", "good 👍", "awesome 😎"},
	{"sad 😞", "bad 👎", "terrible 😩"},
}

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("Please input your name. And you mood (positive or negative)")
		return
	}

	name, mood := os.Args[1], os.Args[2]

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var index int
	if mood != "positive" {
		index = 1
	}

	log.Printf("%s is %s.", name, moods[index][n])
}
