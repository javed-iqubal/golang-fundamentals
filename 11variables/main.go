package main

import "fmt"

const (
	initcount = 10
	lastcount = 100
)

func main() {

	const pi = 3.1415926

	fmt.Println("PI: ", pi)

	// const can't be re-initalized
	// pi = 12

	fmt.Println(initcount, lastcount)
}
