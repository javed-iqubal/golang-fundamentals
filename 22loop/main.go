package main

import "fmt"

// loop:  break & continue

func main() {

	for i := 1; i < 10; i++ {

		if i == 5 {
			continue
		}
		if i == 9 {
			break
		}

		fmt.Println(i)

	}
}
