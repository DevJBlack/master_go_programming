package main

import (
	"fmt"
	"os"
	"strconv"
)

func myFunc(n string) int {
	s, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("Cannot convert %q to an int", n)
		os.Exit(1)
	}

	ss, _ := strconv.Atoi(n + n)
	sss, _ := strconv.Atoi(n + n + n)
	return s + ss + sss
}

func main() {
	sum := myFunc("5")
	fmt.Printf("The sum is: %d", sum)
}
