package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bs := []byte("The Gopher is the golang mascot")
	err := ioutil.WriteFile("info.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// b, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Info in file: %s", b)

	file, err := ioutil.ReadFile("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Here is the file: %s", file)

	// ###########################################

	// ns := []byte("Hello Golang World, this is amazing and very cool!!")
	// err := ioutil.WriteFile("info.txt", ns, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // defer closing the file
	// defer file.Close()

	// // ioutil.ReadAll() reads from the file until an error or EOF and returns the data it read
	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Data as string: %s\n", data)

	// My solution to reading from a file with ioutill
	// With ioutil.ReadFile, you can grab the file you want and read it, with out opening it and closing it.
	// file, err := ioutil.ReadFile("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Inside the file: %s", file)

}
