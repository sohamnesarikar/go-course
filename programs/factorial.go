package main

import "fmt"

func factorial(){
	var n int
	var fact int = 1

	fmt.Println("Enter a number: ")
	fmt.Scan(&n)

	for i:=1; i<=n; i++{
		fact=fact*i
	}

	fmt.Printf("Factorial of %d = %d\n", n, fact)
}