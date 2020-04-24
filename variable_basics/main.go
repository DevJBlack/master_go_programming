package main

import "fmt"

func main() {
	var age int = 10
	newAge := 27
	fmt.Println(newAge)
	fmt.Println(age)

	var name = "Devin"
	// lName := "Black"
	// fmt.Println("Your name is:", name, lName)
	// _ Blank identifier
	_ = name
	s := "learning golang!"
	fmt.Println(s)
	name = "blue"
	fmt.Println(name)

	float := 5.87
	fmt.Println(float)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j //Swapping variables

	fmt.Println(i, j)

	sum := 5 + 2.5
	fmt.Println(sum)
}
