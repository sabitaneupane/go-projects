package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Input please:")
		return
	}

	arg := os.Args[1]

	var richter string

	switch arg {
	case "micro":
		richter = "less than 2.0"
	case "very minor":
		richter = "2.0-2.9"
	case "minor":
		richter = "3-3.9"
	case "light":
		richter = "4-4.9"
	case "moderate":
		richter = "5-5.9"
	case "strong":
		richter = "6-6.9"
	case "major":
		richter = "7-7.9"
	case "great":
		richter = "8-9.9"
	case "massive":
		richter = "10+"
	default:
		richter = "unknown"
	}

	fmt.Printf("%s's richter scale is %s \n", arg, richter)
}
