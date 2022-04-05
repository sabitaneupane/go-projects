package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usage = `Please submit [operator] [size]
CMD: go run main.go [operator] [size]
`
	validOperator = "+-*/"
)

func main() {
	args := os.Args[1:]

	switch l := len(args); {
	case l == 1:
		fmt.Println("Size is missing")
		return
	case l < 1:
		fmt.Println(usage)
		return
	}

	operator := args[0]

	if strings.IndexAny(operator, validOperator) == -1 {
		fmt.Printf("Invalid Operator: %q \n", operator)
		fmt.Printf("Valid Operator are: %q \n", validOperator)
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}
	fmt.Printf("%5s", operator)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var result int

			switch operator {
			case "+":
				result = i + j
			case "-":
				result = i - j
			case "*":
				result = i * j
			case "/":
				if j != 0 {
					result = i / j
				}
			}

			fmt.Printf("%5d", result)
		}
		fmt.Println()
	}

}
