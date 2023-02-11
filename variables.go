package main

import "fmt"

func main() {
	var text string = "Hello, "
	var name string = "World!"

	var date int = 11
	var month = "February"
	var year = 2023

	fmt.Println(text + name)
	fmt.Println("Today is", date, month, year)
}