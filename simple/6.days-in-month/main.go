package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires age")
		return
	}

	arg := os.Args[1]

	month := strings.ToLower(arg)

	year := time.Now().Year()
	isLeapYear := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days := 0

	if month == "january" ||
		month == "march" ||
		month == "may" ||
		month == "july" ||
		month == "august" ||
		month == "october" ||
		month == "december" {
		days = 31

	} else if month == "february" {
		if isLeapYear {
			days = 29
		} else {
			days = 28
		}

	} else if month == "april" ||
		month == "june" ||
		month == "september" ||
		month == "november" {
		days = 30
	} else {
		days = 0
	}

	fmt.Printf("%q has %d days \n", month, days)
}
