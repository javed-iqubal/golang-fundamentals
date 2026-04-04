package main

import "fmt"

func main() {

	fmt.Println("Start")
	fmt.Println("Before defer statement - 1")
	fmt.Println("Before defer statement - 2")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	defer fmt.Println("Fourth defer")
	fmt.Println("After defer statement - 1")
	fmt.Println("After defer statement - 2")
	fmt.Println("End")
}
