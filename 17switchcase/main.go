package main

import "fmt"

func main() {

	age := 16

	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 30:
		fmt.Println("Youth")
	case age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior citizen")
	}
}
