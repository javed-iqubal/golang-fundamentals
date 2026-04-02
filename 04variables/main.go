package main

import "fmt"

func main() {
	fmt.Println("Block level variable declaration")

	// with type
	var (
		name   string  = "Javed Iqubal"
		age    int     = 40
		salary float64 = 75000
		status bool    = true
	)

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Salary: ", salary)
	fmt.Println("Status: ", status)

	fmt.Println()
	// With type inference
	var (
		empName   = "Liza Iqubal"
		empAge    = 20
		empSalary = 25000
		empStatus = true
	)

	fmt.Println("Name: ", empName)
	fmt.Println("Age: ", empAge)
	fmt.Println("Salary: ", empSalary)
	fmt.Println("Status: ", empStatus)
}
