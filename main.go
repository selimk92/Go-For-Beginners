package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	fmt.Println(a)
	b := a[:len(a)-1]
	fmt.Println(b)

}
