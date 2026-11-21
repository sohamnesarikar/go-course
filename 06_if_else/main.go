package main

import "fmt"

func main() {
	// Simple if - else
	age := 20

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}

	// if -  else if - else
	var value int

	fmt.Print("Enter a value: ")
	fmt.Scan(&value)

	if value > 0 {
		fmt.Println("Positive")
	} else if value < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// If with Short Statement
	if num := 13; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
