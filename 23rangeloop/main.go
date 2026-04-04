package main

import "fmt"

// range loop with strings
/*
	for index, value :=range collection/string/arrays/map
*/
func main() {

	str := "Go"

	for index, ch := range str {
		fmt.Println("index: ", index, "value: ", ch)
	}

	fmt.Println()

	str2 := "Python"

	for index, ch := range str2 {
		fmt.Println("index: ", index, "value: ", string(ch))
	}
}
