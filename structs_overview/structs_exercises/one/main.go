package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	colors []string
}

func main() {
	me := person{
		name:   "devin",
		age:    24,
		colors: []string{"blue", "gray", "black", "red"},
	}

	you := person{
		name:   "Austin",
		age:    34,
		colors: []string{"red", "marone", "purple"},
	}

	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)
}
