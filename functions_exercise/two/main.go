package main

import "fmt"

func f1(num uint) (int, int) {
	fact := 1
	sum := 0

	for i := 1; i <= int(num); i++ {
		fact *= i
		sum += i
	}
	return fact, sum
}

func main() {
	fact, sum := f1(5)
	fmt.Printf("Fact: %d, Sum: %d", fact, sum)
}
