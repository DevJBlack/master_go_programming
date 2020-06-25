package main

import "fmt"

func main() {

	x := 10.10
	// 1. Do get the address you need to use the & sign
	fmt.Printf("The address of x is %p\n", &x)

	// 2. This assigns the address to ptr
	ptr := &x

	// 3. We are printing the type which is a float64 and the address
	fmt.Printf("ptr type: %T, ptr value: %v\n", ptr, ptr)

	// 4. First one is the address because we are using &ptr and the second one is the value of x because of the * with the * we are getting 10.10
	fmt.Printf("The address of ptr: %p\n", &ptr)
	fmt.Printf("The value of x through the pointer:%f\n", *ptr)

}
