package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 0:
		fmt.Println("Please give me up to 5 numbers.")
		return
	case l > 5:
		fmt.Println("Sorry. Go arrays are fixed.",
			"So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	var arr [5]float64

	// fill the array with the numbers
	for i, v := range args {
		n, err := strconv.ParseFloat(v, 64)
		if err != nil {
			// skip if it's not a valid number
			continue
		}

		arr[i] = n
	}

	for range arr {
		for i, n := range arr {
			if i < len(arr)-1 && n > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
