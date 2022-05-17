package main

type IntegerDigits [5]string

var zero = IntegerDigits{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = IntegerDigits{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = IntegerDigits{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = IntegerDigits{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = IntegerDigits{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = IntegerDigits{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = IntegerDigits{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = IntegerDigits{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = IntegerDigits{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = IntegerDigits{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon = IntegerDigits{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var digits = [...]IntegerDigits{
	zero, one, two, three, four, five, six, seven, eight, nine,
}
