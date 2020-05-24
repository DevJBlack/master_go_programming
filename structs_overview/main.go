package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine", "Dante Aligehri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespear", 1606

	fmt.Println("Book1", title1, author1, year1)
	fmt.Println("Book2", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		year          int
		title, author string
	}

	// Not recommended could lead to data corruption
	myBook := book{
		"Dante",
		"The Divine",
		1320,
	}
	fmt.Printf("%#v,\n ", myBook)

	// This is the best way
	bestBook := book{
		author: "George",
		year:   2011,
		title:  "Animal",
	}

	fmt.Printf("%#v\n", bestBook)

	aBook := book{title: "Random Title"}
	fmt.Printf("%#v", aBook)

}
