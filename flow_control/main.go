package main

import "fmt"

func main() {

	// if condition_that_evaluates_to_boolean{
	//      perform action1
	// }else if condition_that_evaluates_to_boolean{
	//      perform action2
	// }else{
	//      perform action3
	// }

	price, inStock := 100, true

	if price > 100 {
		fmt.Println("Too Expensive!")
	} else {
		fmt.Println("Perfect price!")
	}
	// _ = inStock
	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	// if price {
	// 	fmt.Println("Something")
	// }

	if price < 100 {
		fmt.Println("Just right")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Don't buy")
	}

	age := 101

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years!\n", 18-age)
	} else if age == 18 {
		fmt.Println("This is your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Println("Please vote!")
	} else {
		fmt.Println("Invalid age")
	}
}
