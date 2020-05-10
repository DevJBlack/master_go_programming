package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	names := []string{"Devin", "Cory", "James"}

	for _, i := range names {
		fmt.Printf("the type is %T, here are the names %v\n", i, i)
	}

	mySlice := []float64{1.2, 5.6}

	// ERROR -> index out of range [2] with length 2
	// mySlice[2] = 6
	mySlice[1] = 6

	// ERROR -> cannot use a (type int) as type float64 in assignment
	// a := 10
	a := 10.
	mySlice[0] = a

	// ERROR -> index out of range [3] with length 2
	// mySlice[3] = 10.10
	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

	// 1. Declare a slice called nums with 3 float64 numbers.
	nums := []float64{1.1, 2.2, 3.3}

	// 2. Append the value 10.1 to the slice
	nums = append(nums, 10.1)

	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	nums = append(nums, 4.1, 5.5, 6.6)

	// 4. Print out the slice
	fmt.Println(nums)

	// 5. Declare a slice called n with 2 float64 values
	n := []float64{10.10, 20.20}

	// 6. Append n to nums
	nums = append(nums, n...)

	// 7. Print out the slice
	fmt.Println(nums)

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	coolNums := []int{0, 11, 2, 8, 9, 1, 11}
	summary := 0

	for _, v := range coolNums[1 : len(coolNums)-1] {
		fmt.Println(v)
		summary += v
	}
	fmt.Println("Sum:", summary)

	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends) //copy(dst, src)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

	dreamCars := []string{"Volvo", "Mercadecs", "Range Rover"}
	newDreamCars := []string{}

	newDreamCars = append(newDreamCars, dreamCars...) // append(empty or new slice, current slice with data then add ...)

	newDreamCars[0] = "Lambo"

	fmt.Println(newDreamCars)

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Printf("%#v\n", newYears)

}
