package main

import "fmt"

func main() {

	// with type
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple variable (a, b ,c): ", a, b, c)
	// with type inference
	var x, y, z = 2, 4, 6
	fmt.Println("Multiple variable (x, y, z): ", x, y, z)

	// with type inference
	var empName, age, isSalaried = "Javed Iqubal", 40, true
	fmt.Println("Employee name: ", empName)
	fmt.Println("Age: ", age)
	fmt.Println("Salaried employee: ", isSalaried)

}
