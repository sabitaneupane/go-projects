package main

import (
	"fmt"
	"os"
)

func main() {
	// words := strings.Fields(
	// 	"Lazy cat jumps again and again and again",
	// )

	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }

	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)

	}

}
