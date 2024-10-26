package main

import "fmt"

func main() {
	i := 5
	switch {
	case i <= 10:
		fmt.Println("one")
	case i <= 20:
		fmt.Println("two")
	case i == 5:
		fmt.Println("three")
	default:
		fmt.Println("default")

	}
}
