package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------------")

	for i := 5; i >= 1; i-- {
		fmt.Println(i)
	}

	fmt.Println("----------------")

	// Even Numbers Print Karna
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("----------------")

	// While loop in go
	num := 1

	for num <= 5 {
		fmt.Println(num)
		num++
	}

	fmt.Println("----------------")

	// Infinite loop
	// for {
	//     fmt.Println("Hello Golang")
	// }

	// 	Multiplication Table
	// User se number input lo aur uska 1 se 10 tak table print karo.
	var userInput int

	fmt.Print("Enter number: ")
	fmt.Scan(&userInput)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", userInput, i, userInput*i)
	}

	// Task 4: Sum of Numbers
	// 1 se 100 tak sabhi numbers ka sum calculate karo.
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)

	// Task 5: Pattern Printing
	// *
	// * *
	// * * *
	// * * * *
	// * * * * *

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
