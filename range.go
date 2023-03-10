package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}

	for num := range nums {
		fmt.Println(num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
}