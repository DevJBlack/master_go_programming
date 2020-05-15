package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string = "devin"
	country := "USA"

	fmt.Printf("Your Name: %s\n", name)
	fmt.Printf("Here is my Country: %s\n", country)

	hello := "He says: \"Hello\" "
	fmt.Println(hello)

	c := "C:\\Users\\a.txt"
	fmt.Println(c)

	r := 'ă'                     // declaring a rune
	fmt.Printf("r type:%T\n", r) // rune is alias to int32

	s1, s2 := "ma", "m" // declaring 2 strings

	// concatenating strings
	s := s1 + s2 + string(r)   // converting rune to string (expliction conversion is required)
	fmt.Printf("s is %s\n", s) // => s is mamă

	s111 := "țară means country in Romanian"

	// iterating over the string and print out byte by byte
	fmt.Printf("Bytes in string: ")
	for i := 0; i < len(s111); i++ {

		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	fmt.Println()

	s11 := "Go is cool!"

	x := s1[0]
	fmt.Println(x)

	// ERROR -> cannot assign to s1[0]
	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s11))

	ss := "你好 Go!"

	// converting string to byte slice
	bbb := []byte(ss)

	// printing out the byte slice
	fmt.Printf("%#v\n", bbb)

	// iterating over the byte slice and printing out each index and byte in the byte slice
	for i, v := range bbb {
		fmt.Printf("%d -> %d\n", i, v)
	}

	sa := "你好 Go!"

	// converting string to rune slice
	ra := []rune(sa)

	// printing out the rune slice
	fmt.Printf("%#v\n", ra)

	// iterating over the rune slice and printing out each index and rune in the rune slice
	for i, v := range ra {
		fmt.Printf("%d -> %c\n", i, v)
	}

}
