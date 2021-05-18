package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("lazy cat jumps again and again and again")

	for i := 0; i < len(words); i++ {
		fmt.Printf("#%-2d: %q\n", i+1, words[i])
	}
	fmt.Println("with range ------")
	for _, v := range words[3:] {
		fmt.Printf("%q\n", v)
	}
}
