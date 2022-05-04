package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs := os.Args[1:]

main:
	for _, arg := range inputs {
		n, err := strconv.Atoi(arg)

		if err != nil {
			continue
		}

		switch {
		case n <= 1:
			continue
		case n == 2 || n == 3:
			fmt.Println(n)
			continue
		}

		for i := 2; i <= n; i++ {
			if n%i == 0 {
				continue main
			} else {
				fmt.Println(n)
				break
			}
		}
	}
}
