package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for num := range numbers {
		if num%2 == 0 {
			fmt.Println(num, "is Even")
		} else {
			fmt.Println(num, "is Odd")
		}
	}

	devin := "Devin Black"

	fmt.Println(devin)
}
