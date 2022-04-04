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
	username1, password1 = "sabita", "test123"
	username2, password2 = "jeevan", "hello00"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != username1 && u != username2 {
		fmt.Println(errUser, u)
	} else if u == username1 && p == password1 {
		fmt.Println(accessOK, u)
	} else if u == username2 && p == password2 {
		fmt.Println(accessOK, u)
	} else {
		fmt.Println(errPass)
	}

}
