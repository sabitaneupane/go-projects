package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires age")
		return
	}

	agr := os.Args[1]
	age, err := strconv.ParseInt(agr, 0, 64)

	if err != nil || age < 0 {
		fmt.Printf("Wrong age: %q \n", agr)
		return
	} else if age > 17 {
		fmt.Println("R-Rated")
		return
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
		return
	} else if age < 13 {
		fmt.Println("PG-RATED")
		return
	}
}
