package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["rank-1"] = "Lionel Messi"
	m["rank-2"] = "Cristiano Ronaldo"

	fmt.Println(m)
	fmt.Println(m["rank-1"])
	fmt.Println(len(m))

	delete (m, "rank-2")

	m["rank-2"] = "Neymar Jr"

	fmt.Println(m)
}