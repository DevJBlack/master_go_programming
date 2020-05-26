package main

import (
	"fmt"
)

type grades struct {
	grade  int
	course string
}

type person struct {
	name   string
	age    int
	colors []string
	school grades
}

func main() {
	me := person{
		name:   "Marius",
		age:    30,
		colors: []string{"red", "yellow"},
		school: grades{
			course: "Python",
			grade:  72,
		},
	}
	you := person{
		name:   "Maria",
		age:    22,
		colors: []string{"pink", "blue"},
		school: grades{
			course: "app development",
			grade:  100,
		},
	}

	me.school.grade = 100
	me.school.course = "Golang"

	fmt.Printf("%+v\n", me)
	fmt.Printf("%v\n", you)

}
