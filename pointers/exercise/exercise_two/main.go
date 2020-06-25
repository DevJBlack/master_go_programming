package main

import "fmt"

func main() {
	// Initialized x and y to 10 and 2
	x, y := 10, 2
	// Assigned ptrx and ptry to the addressess to x and y by using the & sign
	ptrx, ptry := &x, &y
	// Next we needed to get the pointer of x and y, the way to do that is by using the * sign to ptrx and ptry, if you try and do *x/*y you get an error of invalid indirect of x and y. It needs to already have the address assigned to have the pointer point to.
	z := *ptrx / *ptry
	fmt.Printf("z is %d\n", z)

}
