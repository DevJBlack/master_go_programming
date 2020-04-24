package main

import (
	"fmt"
)

func main() {
	a, z := 4, 6
	fmt.Println("sum:", a+z, "Mean Value:", (a+z)/2)

	fmt.Printf("Your age is %d\n", a+z)

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

	// VERBS:
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point

	a, b, c := 10, 15.5, "Gophers"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c) // => a is 10, b is 15.500000, c is Gophers
	fmt.Printf("%q\n", c)                               // => "Gophers"
	fmt.Printf("%v\n", grades)                          // => [10 20 30]
	fmt.Printf("%#v\n", grades)                         // => b is of type float64 and grades is of type []int
	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	// => b is of type float64 and grades is of type []int
	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xc000016128
	fmt.Printf("%c and %c\n", 100, 51011)    // =>  d and 읃  (runes for code points 101 and 51011)

	const newpi float64 = 3.14159265359
	fmt.Printf("pi is %.4f\n", newpi) // => formatting with 4 decimal points

	// %b -> base 2
	// %x -> base 16
	fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
	fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65

	// fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
	s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
	fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers
}
