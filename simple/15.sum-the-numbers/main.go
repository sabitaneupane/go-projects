package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage = `Please submit [operator] [size]
CMD: go run main.go [operator] [size]
`
	validOperator = "+-*/"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: [number]")
		return
	}

	args := os.Args[1]

	number, err := strconv.Atoi(args)

	if err != nil || number < 0 {
		fmt.Println("Wrong Numbers")
		return
	}

	var result int
	for i := 0; i <= number; i++ {
		result = result + i
	}

	fmt.Printf("Total sum upto %d is %d \n", number, result)
}
