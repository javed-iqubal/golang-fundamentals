package main

import "fmt"

// package level variable

func main() {
	// function lavel variable
	var name string = "Javed Iqubal"
	var age uint16 = 40
	var salary float64 = 50000
	var status bool = true

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Salary: ", salary)
	fmt.Println("Status: ", status)
	fmt.Println()

	// type inference: The type of variable is understood based on literal(value)

	var fname = "Javed"
	var lname = "Iqubal"
	var yearOfExperience = 10
	var income = 60000.00
	var isMarriad = true

	fmt.Println("Full Name: ", fname, lname)
	fmt.Println("Year of experience: ", yearOfExperience)
	fmt.Println("Income: ", income)
	fmt.Println("Married: ", isMarriad)

	fmt.Println()
	// type inference with shortcut: The type of variable is understood based on literal(value)

	var fullName = "Javed Iqubal"
	var totalExperience = 9
	var totalIncome = 75000.00
	var isGraduate = true

	fmt.Println("Full Name: ", fullName)
	fmt.Println("Total years of experience: ", totalExperience)
	fmt.Println("Total income: ", totalIncome)
	fmt.Println("Graduate: ", isGraduate)

}
