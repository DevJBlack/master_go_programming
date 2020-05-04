package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python": //values must be comparable (compare string to string)
		fmt.Println("You are learning Python! You don't use { } but indentation !! ")
		// an implicit break is added here
	case "Go", "golang": //compare language with "Go" OR "golang"
		fmt.Println("Good, Go for Go!. You are using {}!")
	default:
		// the default clause the equivalent of the else clause of an if statement
		// and gets executed if no testing condition is true.
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Eve Number!")
	case n%2 != 0:
		fmt.Println("Odd Number!")
	default:
		fmt.Println("Never here!")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)
	switch {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good AfterNoon")
	default:
		fmt.Println("Good Evening")
	}

	//** Switch simple statement **//

	// Syntax: statement (n:=10), semicolon and a switch condition
	//(true in this case, we are comparing boolean expressions that return true)
	// we can remove the word "true" because it's the default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
