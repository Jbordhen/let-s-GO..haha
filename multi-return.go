package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := getRandomNumbers()

	fmt.Println(a)
	fmt.Println(b)
}

func getRandomNumbers() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}