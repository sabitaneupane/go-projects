package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

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

	colon := IntegerDigits{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	digits := [...]IntegerDigits{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear()

	for {
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]IntegerDigits{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		for line := range clock[0] {
			for digit := range clock {
				fmt.Print(clock[digit][line], "  ")
			}
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}
