package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int
	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	num := [2]int{10, 27}
	fmt.Printf("%#v\n", num)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [3]string{"Devin", "James", "Cory"}
	fmt.Printf("%#v\n", a3)

	a4 := [3]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 100, 123}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3,
		4,
		5,
	}
	fmt.Println(a6)

	newNum := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", newNum)

	newNum[0] = 101
	fmt.Printf("%#v\n", newNum)

	// newNum[5] = 100 //Error

	for i, v := range newNum {
		fmt.Printf("index is %d, and value is %d\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < len(newNum); i++ {
		fmt.Println("Index:", i, "Value", newNum[i])
	}

	fmt.Println(strings.Repeat("#", 20))

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m ", n == m)

	m[0] = -1
	fmt.Println("n is equal to m ", n == m)

	var cities [2]string
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{4.5, 9.7}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{true, false, false, true}
	fmt.Printf("%#v\n", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for index, value := range grades {
		fmt.Println(index, value)
	}

	numArray := []int{30, -1, -6, 90, -6}

	// positiveNum := 0 // This way is possible
	var positiveNum int // This is another way

	for _, v := range numArray {
		if v%2 == 0 && v > 0 {
			positiveNum++
		}
	}
	fmt.Println("No. of positive numbers in numArray:", positiveNum)
}
