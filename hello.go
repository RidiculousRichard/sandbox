package main

import (
	"fmt"
)

func main() {
	printHelloWorld()
}

func printHelloWorld() {
	fmt.Println("Hello world")
	nomber := 42
	fmt.Println("The answer to life, the universe and everything is ....", nomber)
	nomber = 41
}
