package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	//tijjj := true
	//fmt.Println(-200, -100, 0, 50, 100, tijjj)
	var x int = math.MaxUint8
	fmt.Println(x)
	msg := os.Args[1]
	l := len(msg)
	rep := strings.Repeat("!", l)
	s := rep + msg + rep
	s = strings.ToUpper(s)
	fmt.Println(s)
	fmt.Printf("%08b", 2)
}
