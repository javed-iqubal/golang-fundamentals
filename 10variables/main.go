package main

import "fmt"

// custom type
type Salary float32

type Celcius float64

func main() {

	//var employeeSalary float32 = 10000

	var employeeSalary Salary = 20000
	var wages Salary = 10000

	fmt.Println(employeeSalary, wages)

	var temp Celcius = 100
	fmt.Println(temp)

}
