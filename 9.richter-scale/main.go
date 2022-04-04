package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Give me the magitude of the earthquake.")
		return
	}

	arg := os.Args[1]
	richter, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var msg string

	switch r := richter; {
	case r < 2:
		msg = "micro"
	case r >= 2 && r < 3:
		msg = "very minor"
	case r >= 3 && r < 4:
		msg = "minor"
	case r >= 4 && r < 5:
		msg = "light"
	case r >= 5 && r < 6:
		msg = "moderate"
	case r >= 6 && r < 7:
		msg = "strong"
	case r >= 7 && r < 8:
		msg = "major"
	case r >= 8 && r < 10:
		msg = "great"
	default:
		msg = "massive"
	}

	fmt.Printf("%.2f is %s\n", richter, msg)
}
