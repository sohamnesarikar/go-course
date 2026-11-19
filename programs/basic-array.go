package main

import "fmt"

func arrayBasics(){
	var numbers [3]int
	numbers[0] = 10
	numbers[1] =20
	numbers[2]=30

	fmt.Println(len(numbers))
	fmt.Println(numbers)

	// declaration and initialization
	values:= [...]string{"Rohit", "Virat", "Dhoni"}

	for i:=0; i<len(values); i++{
		fmt.Println(values[i])
	}

	// range
	for index, value:= range numbers {
		fmt.Println(index, value)
	}

	var marks []int
	fmt.Println(marks) // empty array
}