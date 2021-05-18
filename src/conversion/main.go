package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please tell me value in feet!")
		return
	}
	const (
		fahrenheitCst = 1.8
		degreCst      = 32
	)
	arg := os.Args[1]
	celsus, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("Error occured, %q is not a number.", arg)
		return
	}
	fahrenheit := celsus*fahrenheitCst + degreCst // 1.8 + 32
	fmt.Printf("%g°C is %g°F.\n", celsus, fahrenheit)
	str := "Dramé"
	fmt.Println("bytes:", len(str))
	fmt.Println("taille:", utf8.RuneCountInString(str))
	fmt.Println(math.Round(6. / 7.))
}
