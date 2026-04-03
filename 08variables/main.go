package main

import "fmt"

func main() {

	var name string = "Javed Iqubal"
	fmt.Println("Name: ", name)

	var htmlDoc string = `
	<html>
	<head><meta charset="utf-8"/>
	<title>Go Language Fundamentals</title>
	</head>
	<body>
	<h1>Go Language</h1>
	</body>
	</html>	
	`
	fmt.Println(htmlDoc)
}
