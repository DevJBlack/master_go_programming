package main

import (
	"fmt"
	"math"
)

func main() {
	x := cube(3)
	fmt.Println(x)
	// fmt.Println(cube(3))
}

func cube(num float64) float64 {
	return math.Pow(num, 3)
	// Or you can do
	// return num * num * num
}
