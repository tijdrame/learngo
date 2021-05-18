package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	entries := os.Args[1:]
	words := strings.Fields("lazy cat jumps again and again and again")

queries:
	for _, it := range entries {

		for i, v := range words {
			if strings.EqualFold(v, it) {
				fmt.Printf("#%-2d: %q\n", i+1, v)
				break queries
			}
		}
	}
}
