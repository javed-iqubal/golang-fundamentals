package main

import (
	"fmt"
)

func main() {

	language := "java"

	switch language {
	case "go":
		fmt.Println("Go language")
	case "python":
		fmt.Println("Python")
	case "java":
		fmt.Println("Java")
	case "c":
		fmt.Println("C")
	case "cpp":
		fmt.Println("C++")
	default:
		fmt.Println("Unknown lnguage")
	}
}
