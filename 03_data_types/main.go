package main

import "fmt"

func main() {
	// Basic types
	var name string = "Soham"

	var age int = 28

	var isMarried bool = false

	var price float64 = 200.67

	fmt.Println(name, age, isMarried, price)

	// byte
	var char byte = 'A'
	fmt.Println(char)

	// rune
	var emoji rune = '😎'
	fmt.Println(emoji)         // unicode
	fmt.Printf("%c \n", emoji) // actual emoji
}
