package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	var name string
	fmt.Println(name)
	name = "Chirag"
	fmt.Println(name)
}

var name = "Chirag"

func printName() {
	fmt.Println(name)
}
