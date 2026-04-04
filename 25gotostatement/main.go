package main

import "fmt"

// goto used as loop

func main() {

	i := 0
loop:

	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}
}
