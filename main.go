package main

import "fmt"

func main() {
loop:
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break loop
			}

		}

	}
}
