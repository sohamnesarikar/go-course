package main

import "fmt"

func multiplicationTable(){
	var n int

	fmt.Println("Enter number: ")
	fmt.Scan(&n)

	for i:=1; i<=10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n * i)
	}
}