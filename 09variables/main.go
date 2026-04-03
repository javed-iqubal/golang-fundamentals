package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "Javed Iqubal"

	length := len(name)
	fmt.Println(length)

	fmt.Println(strings.ToUpper(name))
	name = strings.ToLower(name)
	words := strings.Split(name, " ")
	fmt.Println(words)

}
