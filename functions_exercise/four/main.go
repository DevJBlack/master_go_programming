package main

import (
	"fmt"
)

func all(sum ...int) int {
	value := 0

	for _, v := range sum {
		value += v
	}

	return (value)
}

func main() {
	a := all(1, 2, 3, 4, 5)
	fmt.Println(a)
}
