package main

import "fmt"

func main() {
	day := 9

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Week days")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}
