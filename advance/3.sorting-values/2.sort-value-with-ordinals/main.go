package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Send me some itams and I will sort them")
		return
	}

	args := os.Args[1:]
	sort.Strings(args)

	var n []byte

	for i, v := range args {
		n = strconv.AppendInt(n, int64(i+1), 10)
		n = append(n, '.', ' ')
		n = append(n, v...)
		n = append(n, '\n')
	}

	err := ioutil.WriteFile("out.txt", n, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
}
