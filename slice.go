package main

import "fmt"

func main() {
	x:=make([]int, 0)

	fmt.Println(x)

	fmt.Println(len(x))

	x = append(x, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52)

	fmt.Println(x)
}