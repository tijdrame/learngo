package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Give a number & type of calculation")
		return
	}
	s := os.Args[1]
	num, _ := strconv.ParseInt(s, 0, 0)
	numI := int(num)
	fmt.Printf("num %d\n", numI)
	typ, signe := os.Args[2], ""
	switch typ {
	case "*":
		signe = "*"
	case "/":
		signe = "/"
	case "+":
		signe = "+"
	case "-":
		signe = "-"
	}
	fmt.Printf("%5s", typ)
	for i := 1; i <= numI; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 1; i <= numI; i++ {
		fmt.Printf("%5d", i)
		for j := 1; j <= numI; j++ {
			if signe == "*" {
				fmt.Printf("%5d", i*j)
			} else if signe == "/" {
				fmt.Printf("%5d", i/j)
			} else if signe == "+" {
				fmt.Printf("%5d", i+j)
			} else {
				fmt.Printf("%5d", i-j)
			}

		}
		fmt.Println()
	}
	/*sum := 0
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("Sum= ", sum)*/
}
