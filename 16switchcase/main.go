package main

import "fmt"

func main() {

	switch num := 10; num {
	case 20:
		fmt.Println("Twenty")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("default")
	}
}
