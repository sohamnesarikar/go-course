package main

import "fmt"

func main() {
	var num int
	correctGuess := 17
	n := 30

	for num != correctGuess {
		fmt.Printf("Guess the number between 1 & %d: ", n)
		_, err := fmt.Scan(&num)

		if err != nil {
			fmt.Println("Please enter valid number...")
			continue
		}

		if num == correctGuess {
			fmt.Printf("Congratulations Thats right correct number is %d\n", num)
		} else if num > correctGuess {
			fmt.Println("Wrong! Too high...")
		} else {
			fmt.Println("Wrong! Too Low...")
		}
	}
}
