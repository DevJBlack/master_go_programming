package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	devin := []string{"Devin James Black"}

	for _, name := range devin {
		fmt.Printf("%q", strings.Split(name, ""))
	}

	for n := 1; n <= 3; n++ {
		fmt.Println("Ho! ")
	}
	fmt.Println("Merry Christmas!")

	// has the same effect as a while loop in other languages
	// there is no while loop in Go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	//** CONTINUE STATEMENT **//

	// It works just the same as in C,  Java or Python.
	// The continue statement rejects all the remaining statements in the current iteration of the loop
	// and moves the control back to the top of the loop.

	// printing even numbers less than or equal to 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skipping the remaining code in this iteration
		}
		fmt.Println(i)
	}

	// **BREAK STATEMENT **//

	// It is used to terminate the innermost for or switch statement.
	// It works just the same as in C,  Java or Python.

	// finding 10 numbers divisible by 13
	count := 0
	for i := 0; true; i++ { // we use true here to run an infinite loop, we can we replace true with 1000 and it will still only print 10 numbers
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 { //if 10 numbers were found, break!
			break //it breaks the current loop (inner loop if there are more loops)
		}
	}

	// the break statement is not terminating the program entirely;
	fmt.Println("Just a message after the for loop")
}
