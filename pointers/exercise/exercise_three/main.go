package main

import "fmt"

// Arguments that we are taking in are the address from func main &x and &y
// Then in the swap function, we are pointing to those addressess using pointers *a, *b
func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	x, y := 5.5, 8.8
	// This is the exact same thing, but instead of assigning variables  to the addressess of x and y we just put them in directly using &x and &y
	swap(&x, &y)
	// This is way works just fine, we are assigning variables to &x and &y and passing them into the function
	// ptrx := &x
	// ptry := &y
	// swap(ptrx, ptry)

	fmt.Printf("x is %v and y is %v\n", x, y)

}
