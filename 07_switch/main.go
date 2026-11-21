package main

import "fmt"

func main() {
	// Simple Switch
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	default:
		fmt.Println("Invalid day")
	}

	// Multiple Cases
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Working days")

	case 6, 7:
		fmt.Println("Weekdays")

	default:
		fmt.Println("Invalid day")
	}

	// Switch Without Expression
	var num1, num2, num3 int

	fmt.Print("Enter num1: ")
	fmt.Scan(&num1)

	fmt.Print("Enter num2: ")
	fmt.Scan(&num2)

	fmt.Print("Enter num3: ")
	fmt.Scan(&num3)

	switch {
	case num1 >= num2 && num1 >= num3:
		fmt.Println("Num1 is greater", num1)

	case num2 >= num1 && num2 >= num3:
		fmt.Println("Num2 is greater", num2)

	default:
		fmt.Println("Num3 is greater", num3)
	}

	// fallthrough
	number := 1

	switch number {
	case 1:
		fmt.Println("One")
		fallthrough

	case 2:
		fmt.Println("Two")

	default:
		fmt.Println("Other")
	}
}
