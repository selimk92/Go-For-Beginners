package main

import (
	"fmt"
)

func main() {

	a, b := 1, 0
	ans := a / b
	panic(ans)
	fmt.Println(ans)
}
