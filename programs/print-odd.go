package main

import "fmt"

func printOddNumbers(){
	var n int

	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	for i:=1; i<=n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}