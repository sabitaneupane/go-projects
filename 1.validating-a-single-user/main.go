package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied to user:"
	errPass  = "Invalid password"
	accessOK = "Access granted to user:"
)

const (
	username = "sabita"
	password = "test123"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != username {
		fmt.Println(errUser, u)
	} else if p != password {
		fmt.Println(errPass)
	} else {
		fmt.Println(accessOK, u)
	}

}
