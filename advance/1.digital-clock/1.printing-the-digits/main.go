package main

import "fmt"

type IntegerDigits [5]string

func main() {

	zero := IntegerDigits{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := IntegerDigits{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := IntegerDigits{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := IntegerDigits{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := IntegerDigits{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := IntegerDigits{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := IntegerDigits{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := IntegerDigits{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := IntegerDigits{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := IntegerDigits{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := []IntegerDigits{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for _, num := range digits {
		for _, n := range num {
			fmt.Println(n)
		}
		fmt.Println("")
	}
}
