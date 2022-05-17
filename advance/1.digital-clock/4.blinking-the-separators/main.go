package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// for keeping things easy to read and type-safe
	type IntegerDigits [5]string

	// put the digits (IntegerDigitss) into variables
	// using the IntegerDigits array type above
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

	// This array's type is "like": [10][5]string
	//
	// However:
	// + "IntegerDigits" is not equal to [5]string in type-wise.
	// + Because: "IntegerDigits" is a defined type, which is different
	//   from [5]string type.
	// + [5]string is an unnamed type.
	// + IntegerDigits is a named type.
	// + The underlying type of [5]string and IntegerDigits is the same:
	//     [5]string
	digits := [...]IntegerDigits{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	// For Go Playground, do not use this.
	screen.Clear()

	// Go Playground will not run an infinite loop.
	// Loop for example 1000 times instead, like this:
	//   for i := 0; i < 1000; i++ { ... }
	for {
		// For Go Playground, use this instead:
		// fmt.Print("\f")
		screen.MoveTopLeft()

		// get the current hour, minute and second
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// extract the digits: 17 becomes, 1 and 7 respectively
		clock := [...]IntegerDigits{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		// Explanation: clock[0]
		// + Each element of clock has the same length.
		// + So: Getting the length of only one element is OK.
		// + This could be: "zero" or "one" and so on... Instead of: digits[0]
		//
		// The range clause below is ~equal to the following code:
		//   line := 0; line < len(clock[0]); line++
		for line := range clock[0] {
			// Print a line for each IntegerDigits in clock
			for index, digit := range clock {
				// Colon blink on every two seconds.
				// + On each sec divisible by two, prints an empty line
				// + Otherwise: prints the current pixel
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				// Print the next line and,
				// give it enough space for the next IntegerDigits
				fmt.Print(next, "  ")
			}

			// After each line of a IntegerDigits, print a newline
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}
