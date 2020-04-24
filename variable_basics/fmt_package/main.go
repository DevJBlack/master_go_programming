package main

import (
	"fmt"
)

func main() {
	a, b := 4, 6
	fmt.Println("sum:", a+b, "Mean Value:", (a+b)/2)

	fmt.Printf("Your age is %d\n", a+b)

	fmt.Printf("x is %d, y is %f", 5, 6.8)
	fmt.Printf("he says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	fmt.Printf("Radius is %d\n", radius)

	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant %f\n", pi)

	fmt.Printf("The diameter of %s with a radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	//%q for quoted string
	fmt.Printf("This is a %q\n", figure)

	//%v -> replace by anytype
	fmt.Printf("The diameter of %v with a radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	//%T -> type
	fmt.Printf("figure is of type %T \n", figure)
	fmt.Printf("radius is of type %T \n", radius)
	fmt.Printf("pi is of type %T \n", pi)

	// %t -> bool
	closed := true
	fmt.Printf("file closed: %t\n", closed)

	// %b -> base 2
	fmt.Printf("%08b \n", 8)

	// %x
	fmt.Printf("%x \n", 100)

	x := 3.4
	y := 6.9

	fmt.Println(x * y)
	fmt.Printf("x * y = %.3f", x*y)
}
