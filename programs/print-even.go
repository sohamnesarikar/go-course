package main

import "fmt"

func printEvenNumbers(){
	var i int = 1;

	for i<=50 {
		if i%2==0 {
			fmt.Println(i)
		}
		i++
	}
}