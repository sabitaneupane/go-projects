package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide directory paths")
		return
	}

	args := os.Args[1:]

	var dirs []byte

	for _, dir := range args {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')
		fmt.Println(dirs)

		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...)
				dirs = append(dirs, '/', '\n')
			}
		}
		// dirs = append(dirs, '\n')
	}
	err := ioutil.WriteFile("dirs.txt", dirs, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
