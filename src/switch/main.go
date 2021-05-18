package main

import (
	"fmt"
	"time"
)

func main() {
	//t := time.Now()
	//h := t.Hour()
	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("good evening")
	case h >= 12:
		fmt.Println("good afternoon")
	case h >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good night")

	}

	/*i := 105
	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}*/
}
