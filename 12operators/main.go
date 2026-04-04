package main

import "fmt"

func main() {

	// base values

	a, b, c, d := 10, 30, 50, 40
	x, y := true, false
	p, q := 12, 9

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

	// comparision
	fmt.Printf("Comparision: %d == %d -> %t, %d > %d -> %t \n", a, b, a == b, c, d, c > d)

	// Logical
	fmt.Printf("Logical: %t && %t -> %t, %t||%t -> %t, !%t -> %t \n", x, y, x && y, x, y, x || y, x, !x)

	//Bitwise
	fmt.Printf("Bitwise: %b & %b -> %b (%d), %b | %b -> %b (%d) \n", p, b, p&b, p, b, p|b)

	// Shift
	fmt.Printf("Shift: %b << 1 = %b (%d), %b >> 1 = %b (%d) \n", p, p<<1, p<<1, p, p>>1, p>>1)

	// AND NOT
	fmt.Printf("AND NOT: %b &^ %b = %b (%d) \n", p, 8, p&^8, p&^8)
	fmt.Printf("AND NOT: %b &^ %b = %b (%d) \n", q, 1, q&^1, q&^1)

}
