package main

import "fmt"

// Golang doesn't have while loop, so we can use for loop as while loop

/*
	// decalation
	for condition{
		// do something
		// post/pre operations
	}
*/

func main() {

	i := 1

	for i < 10 {
		fmt.Println(i)
		i++
	}

}
