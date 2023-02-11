package main

import "fmt"

func isOdd(n int) bool {
	return n%2 == 1
}

func main() {
	fmt.Println(isOdd(1))
	fmt.Println(isOdd(2))
}