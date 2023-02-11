package main

import "fmt"

func main() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("Back to work again.")
	case "friday":
		fmt.Println("Time to party!")
	default:
		fmt.Println("Back to work")
	}
}
