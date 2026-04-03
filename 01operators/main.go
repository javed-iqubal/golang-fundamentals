package main

import "fmt"

func main() {

	// base values

	a, b, c, d := 10, 30, 20, 50
	//x, y := true, false
	//p := 12

	// Arithmetic

	fmt.Printf("Arithmetic: %d+%d=%d, %d-%d=%d, %d*%d=%d \n", a, b, a+b, a, b, a-b, a, b, a*b)

	// Division & Modulus
	fmt.Printf("Division: %d/%d= %d, %d %% %d=%d \n", b, a, b/a, d, c, d%c)

	// Unary +/-
	fmt.Printf("Unary: %d=%d, -%d=%d \n", a, +a, b, -b)

	// Increment/Decrement (statement only)
	a++
	b++
	fmt.Printf("Increment/Decrement: a++ -> %d b-- -> %d\n", a, b)

	// Assignment (=)
	e := 5
	fmt.Printf("Assignment: e=5 -> %d \n", e)

	// compund
	e += 2
	fmt.Printf("Compound: e += 2 -> %d \n", e)

}
