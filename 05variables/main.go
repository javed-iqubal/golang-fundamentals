package main

import "fmt"

func main() {
	// with type
	var age int = 35
	fmt.Println("Age-1: ", age)
	age = 40
	fmt.Println("Age-2: ", age)

	fmt.Println()
	// with type inference
	var price = 2500.50
	fmt.Println("Price-1: ", price)
	price = 3500.50
	fmt.Println("Price-2: ", price)

	fmt.Println()
	// with type inference with shortcut
	quantity := 250.5
	fmt.Println("Quantity-1: ", quantity)

	quantity = 500
	fmt.Println("Quantity-2: ", quantity)

	fmt.Println()
	var distance uint = 200
	fmt.Println("distance-1: ", distance)

	// cannot use -50 (untyped int constant) as uint value in assignment (overflows)compilerNumericOverflow
	// distance = -50
	// fmt.Println("distance-2: ", distance)

}
