package main

import (
	"fmt"
)

func main() {
	printHelloWorld()
}

func printHelloWorld() {
	fmt.Println("Hello world")
	n := 42
	fmt.Println("The answer to life, the universe and everything is ....", n)
}
