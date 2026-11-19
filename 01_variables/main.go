package main

import "fmt"

func main() {
	// Way 1: variable creation with var keyword
	var firstName string = "Soham"
	var age int = 28
	fmt.Println("name", firstName)
	fmt.Println("age", age)

	// Way 2: var keyword without data type
	var isAdult = true
	fmt.Println("isAdult", isAdult)

	// Way 3: Short variable declaration (:=)
	username := "John"
	password := 23987

	username = "tom12"
	fmt.Println(username, password)

	// Way 4: Multiple variable declaration
	var (
		favLang  string = "Go"
		isFamous bool   = true
	)
	fmt.Println(favLang, isFamous)

	var a, b, c int = 10, 20, 30
	fmt.Println(a, b, c)
}
