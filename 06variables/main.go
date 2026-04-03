package main

import "fmt"

// variable outside function - package level variable

// var age int = 35

// with type inference
var age = 40

// short notation - not allowed out side the function
// age := 38	// syntax error: non-declaration statement outside function body

func main() {
	// variable inside function - local variavle
	var name = "Javed Iqubal"

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
}
