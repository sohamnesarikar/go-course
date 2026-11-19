package main

import "fmt"

func countDigit(){
	var num int
	var count int = 0
	

	fmt.Println("Enter a number: ")
	fmt.Scan(&num)

	for num != 0 {
		count++
		num = num /10
	}

	fmt.Printf("count = %d\n", count)
}