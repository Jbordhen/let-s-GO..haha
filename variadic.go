package main

import (
	"fmt"
)

func main() {
	sum := getSum(1, 2, 3, 4, 5)

	fmt.Println(sum)
}

func getSum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}