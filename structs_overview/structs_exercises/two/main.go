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
	me := person{name: "Connor", age: 24, colors: []string{"blue", "gray", "black"}}
	you := person{name: "Tyler", age: 22, colors: []string{"orange", "red"}}

	// 1.
	me.name = "Devin"

	// 2.
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

	// 3.
	colors = append(colors, "green")
	you.colors = colors

	// 4.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}
}
