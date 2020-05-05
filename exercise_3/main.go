package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		if i%7 == 0 { // if i is divisible by 7
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")

	for i := 1; i <= 50; i++ {
		if i%7 != 0 { // if i is not divisible by 7
			continue
		}
		fmt.Printf("%d ", i)

	}
	fmt.Println("")

	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 { // if i is not divisible by 7
			continue
		}
		fmt.Printf("%d ", i)
		count++

		if count == 3 { // if i've already found 3 numbers, then break
			break
		}

	}
	fmt.Println("")

	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 { // if i is divisible both by 7 and 5
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")

	birthYear, currentYear := 1995, 2020

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	age := 10
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}

}
