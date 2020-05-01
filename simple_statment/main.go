package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The type of i is %T, and the value is %v", i, i)
	}

	if a, err := strconv.Atoi("20"); err == nil {
		fmt.Println("No error, a is:", a)
	} else {
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)/0.621)
		fmt.Printf("%T\n", args[1])
	}

}
