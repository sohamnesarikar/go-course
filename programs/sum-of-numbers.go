package main

import "fmt"

func sumOfNumbers(){
	var sum int = 0;

	for i:=1; i<=100; i++ {
		sum = sum + i
	}

	fmt.Println("Sum:", sum)
}