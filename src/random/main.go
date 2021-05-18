package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Give a number")
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Not a number")
		return
	}
	if guess < 0 {
		fmt.Println("Pick a positive number")
		return
	}
	rand.Seed(time.Now().UnixNano())
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
