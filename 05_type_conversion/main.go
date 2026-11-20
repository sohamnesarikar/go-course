package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to float
	a := 234
	convertToFloat := float64(a)
	fmt.Println("int to float", convertToFloat)
	fmt.Printf("Type: %T\n", convertToFloat)

	// float to int
	b := 345.648
	convertToInt := int(b)
	fmt.Printf("Value: %d, type: %T\n", convertToInt, convertToInt)

	// int to string
	age := 65
	convertToString := strconv.Itoa(age)
	fmt.Printf("Value: %s, type: %T\n", convertToString, convertToString)

	// string to int
	num := "23"
	value, _ := strconv.Atoi(num)
	fmt.Printf("Value: %d, type: %T\n", value, value)
}
