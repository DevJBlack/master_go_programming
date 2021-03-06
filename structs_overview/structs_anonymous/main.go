package main

import "fmt"

func main() {
	// an anonymous struct is a struct with no explicitly defined struct type alias.
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	// =>struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30

	fmt.Printf("Diana's Age: %d\n", diana.age)

	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type employee struct {
		name   string
		salary int
		bool
	}

	e := employee{"John", 40000, false}
	fmt.Printf("%#v\n", e)

	e.bool = true

	fmt.Printf("%#v\n", e)
}
