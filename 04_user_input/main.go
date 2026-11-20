package main

import "fmt"

func main() {
	var a float64
	var b float64

	fmt.Print("Enter num1: ")
	fmt.Scan(&a)

	fmt.Print("Enter num2: ")
	fmt.Scan(&b)

	fmt.Println("--- Calculator Results ---")

	fmt.Printf("Addition: %.2f\n", a+b)
	fmt.Printf("Subtraction: %.2f\n", a-b)
	fmt.Printf("Multiplication: %.2f\n", a*b)
	fmt.Printf("Division: %.2f", a/b)

}
