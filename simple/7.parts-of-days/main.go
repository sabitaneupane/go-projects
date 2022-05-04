package main

import (
	"fmt"
	"time"
)

func main() {
	switch hour := time.Now().Hour(); {
	case hour >= 18:
		fmt.Println("Good Evening")
	case hour >= 12:
		fmt.Println("Good Afternoon")
	case hour >= 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}
}
