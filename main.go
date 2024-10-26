package main

import "fmt"

func main() {
	var i interface{} = 1.5
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")

	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")

	}
}
