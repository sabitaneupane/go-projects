package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Send me some itams and I will sort them")
		return
	}

	args := os.Args[1:]
	sort.Strings(args)

	var n []byte

	for _, v := range args {
		n = append(n, v...)
		n = append(n, '\n')
	}

	err := ioutil.WriteFile("out.txt", n, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
