package main

import (
	"fmt"
	"os"
)

func main() {
	//var name string
	name, name2, name3 := os.Args[1], os.Args[2], os.Args[3]
	fmt.Println("Hello great", name)
	fmt.Println("Hello great", name2)
	fmt.Println("Hello great", name3)
	name, age := "gandalf", 2019
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, you shall pass!")
}
