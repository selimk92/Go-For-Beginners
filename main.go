package main

import "fmt"

func main() {
	number := 50
	guess := 34
	if guess < 1 || guess > 100 {

		fmt.Println("The guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("Guess:", guess, "is below", number)
		}
		if guess > number {
			fmt.Println("Guess:", guess, "is above", number)
		}
		if guess == number {
			fmt.Println("Guess is equal to number")
		}
	}
}
