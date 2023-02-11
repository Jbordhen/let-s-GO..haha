package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("Time to party!")
	default:
		fmt.Println("Back to work")
	}
}
