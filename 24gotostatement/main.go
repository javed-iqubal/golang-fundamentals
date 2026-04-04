package main

import "fmt"

func main() {

	fmt.Println("Start....")
	fmt.Println("Statement 1")
	goto END
	// Unrechable code
	fmt.Println("Statement 2")
	fmt.Println("Statement 3")
	fmt.Println("Statement 4")
	fmt.Println("Statement 5")
END:
	fmt.Println("Statement 6")
	fmt.Println("Statement 7")
	fmt.Println("Statement 8")
	fmt.Println("End")
}
