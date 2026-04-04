package main

import "fmt"

// defer - used to delay the execution of a function until the surrounding function return

func main() {

	fmt.Println("Start")
	fmt.Println("Statement-1")
	fmt.Println("Statement-2")
	defer fmt.Println("delayed function")
	fmt.Println("Statement-3")
	fmt.Println("Statement-4")
	fmt.Println("End")

}
