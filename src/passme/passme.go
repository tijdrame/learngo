package main

import (
	"fmt"
	"os"
)

const (
	username      = "tij"
	password      = "passer"
	u2            = "jack"
	p2            = "1111"
	paramsMiss    = "Give a username and or password"
	wrongUser     = "Username don't match! %q\n"
	wrongPwd      = "Password don't match for %q\n"
	accessGranted = "Access granted!"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(paramsMiss)
		return
	}
	u, p := os.Args[1], os.Args[2]
	if username == u && password == p {

		fmt.Println(accessGranted)

	} else if u2 == u && p2 == p {
		fmt.Println(accessGranted)
	} else {
		fmt.Printf(wrongUser, u)

	}
}
