package main

import (
	f "fmt"
	"path"
	"runtime"
)

/*
main function
*/
func main() {
	name, lastName := "Tidiane", "Dramé"

	f.Println("Hello Gopher!")
	f.Println("other import")
	f.Println(runtime.NumCPU())
	//bye()
	//hey()
	f.Println(name, lastName)
	//var file string
	_, file := path.Split("css/main.css")
	f.Println("file=", file)
	speed := 100
	force := 2.5
	speed = int(float64(speed) * force)
	f.Println(speed)
	//var brand string
	f.Printf("%v\n", speed)
}
